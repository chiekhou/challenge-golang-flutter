import 'dart:convert';
import 'package:flutter/cupertino.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;

import '../models/member_model.dart';

class AuthProvider extends ChangeNotifier{
  final String host = "10.0.2.2";
  final FlutterSecureStorage _storage = FlutterSecureStorage();
  final String _baseUrl = "http://localhost:8080"; // URL de votre backend

  bool get isAuthenticated => _storage.read(key: 'auth_token') != null;
  bool get isAdmin => true;


  //Appel Api register
  Future<bool> Register({
    required String address,
    required String email,
    required String firstname,
    required String lastname,
    required String password,
    required String username,
    required String photo,
  }) async {
    try {
      final response = await http.post(
        Uri.parse('http://$host:8080/Signup'),
        headers: {
          'Content-Type': 'application/json',
        },
        body: jsonEncode({
          'address': address,
          'email': email,
          'first_name': firstname,
          'last_name': lastname,
          'password': password,
          'username': username,
          'photo': photo,
        }),
      );

      print(response.body);
      if(response.statusCode == 201){
        print(response.body);
        return true;
      }else{
        return false;
      }
    }catch(e){
      rethrow;
    }

  }

  //Appel APi pour se loguer
  Future<bool> login(String email, String password) async{
    try{
      final response = await http.post(
          Uri.parse('http://$host:8080/login'),
          headers: {
            'Content-Type': 'application/json'
          },
          body: jsonEncode({
            'email': email,
            'password': password
          })
      );

      print(response.body);

      if(response.statusCode == 200) {
        final responseData = jsonDecode(response.body.toString());
        final token = responseData["token"];

        await _storage.write(key: 'auth_token', value: token);
        return true;
      }else{
        return false;
      }
    }
    catch(e){
      rethrow;
    }
  }

  Future<bool> isLoggedIn() async {
    String? token = await _storage.read(key: 'auth_token');
    return token != null;
  }

  //Appel pour se logout
  Future<void> logout() async {
    try {
      String? token = await _storage.read(key: 'auth_token');
      print('Token récupéré: $token');

      if (token != null) {
        final response = await http.post(
          Uri.parse('http://$host:8080/logout'),
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
        );

        print('Réponse du serveur: ${response.statusCode}');
        print('Corps de la réponse: ${response.body}');

        if (response.statusCode == 200) {
          await _storage.delete(key: 'auth_token');
          print('Déconnexion réussie');
          notifyListeners();
        } else {
          print('Échec de la déconnexion: ${response.body}');
          throw Exception('Failed to logout');
        }
      } else {
        print('Aucun token trouvé');
        throw Exception('No token found');
      }
    } catch (e) {
      print('Erreur de déconnexion: $e');
      rethrow;
    }
  }


  //Profil du user
  Future<Member> Profile()async{
    String? token = await _storage.read(key: 'auth_token');
    if(token != null){
      final response = await http.get(
        Uri.parse('http://$host:8080/profile'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token'
        },
      );
      if(response.statusCode == 200){
        Map<String, dynamic> data = jsonDecode(response.body);
        Map<String, dynamic> userData = data['user'];
        print(response.body);
        print(Member.fromJson(userData));
        return Member.fromJson(userData);
      }else{
        throw Exception('No profile found');
      }
    }else{
      throw Exception('User not logged in');
    }
  }
}
