import 'dart:convert';
import 'package:flutter/cupertino.dart';
import 'package:flutter_app/config/app_config.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;

import '../models/member_model.dart';

class AuthProvider extends ChangeNotifier {
  final FlutterSecureStorage _storage = FlutterSecureStorage();
  final apiAuthority = AppConfig.getApiAuthority();
  final isSecure = AppConfig.isSecure();

  bool get isAuthenticated => _storage.read(key: 'auth_token') != null;

  //Appel Api register
  Future<bool> Register(String address, String email, String firstname,
      String lastname, String password, String username, String photo) async {
    try {
      final url = isSecure
          ? Uri.https(apiAuthority, '/Signup')
          : Uri.http(apiAuthority, '/Signup');
      final response = await http.post(
        url,
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
      if (response.statusCode == 201) {
        print(response.body);
        return true;
      } else {
        return false;
      }
    } catch (e) {
      rethrow;
    }
  }

  //Appel APi pour se loguer
  Future<bool> Login(String email, String password) async {
    try {
      final url = isSecure
          ? Uri.https(apiAuthority, '/login')
          : Uri.http(apiAuthority, '/login');

      final response = await http.post(url,
          headers: {'Content-Type': 'application/json'},
          body: jsonEncode({'email': email, 'password': password}));

      print(response.body);

      if (response.statusCode == 200) {
        final responseData = jsonDecode(response.body.toString());
        final token = responseData["token"];

        await _storage.write(key: 'auth_token', value: token);
        return true;
      } else {
        return false;
      }
    } catch (e) {
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
        final url = isSecure
            ? Uri.https(apiAuthority, '/logout')
            : Uri.http(apiAuthority, '/logout');
        final response = await http.post(
          url,
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
  Future<Member> Profile() async {
    String? token = await _storage.read(key: 'auth_token');
    if (token != null) {
      final url = isSecure
          ? Uri.https(apiAuthority, '/profile')
          : Uri.http(apiAuthority, '/profile');
      final response = await http.get(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token'
        },
      );
      if (response.statusCode == 200) {
        Map<String, dynamic> data = jsonDecode(response.body);
        Map<String, dynamic> userData = data['user'];
        print(response.body);
        print(Member.fromJson(userData));
        return Member.fromJson(userData);
      } else {
        throw Exception('No profile found');
      }
    } else {
      throw Exception('User not logged in');
    }
  }
}
