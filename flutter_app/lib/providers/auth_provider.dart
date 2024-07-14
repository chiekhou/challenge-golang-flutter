import 'dart:convert';
import 'package:flutter/cupertino.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;

class AuthProvider with ChangeNotifier {
  final String host = "10.0.2.2"; // Utilisé pour émulateur Android
  final FlutterSecureStorage _storage = FlutterSecureStorage();
  final String _baseUrl = "http://localhost:8080"; // URL de votre backend

  bool get isAuthenticated => _storage.read(key: 'auth_token') != null;

  bool get isAdmin => false; // Implémentation à adapter selon votre logique

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
        Uri.parse('$_baseUrl/signup'),
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

      if (response.statusCode == 201) {
        print(response.body);
        return true;
      } else {
        print(response.body);
        return false;
      }
    } catch (e) {
      print('Erreur lors de l\'inscription: $e');
      rethrow;
    }
  }

  Future<bool> login(String email, String password) async {
    try {
      final response = await http.post(
        Uri.parse('$_baseUrl/login'),
        headers: {
          'Content-Type': 'application/json',
        },
        body: jsonEncode({
          'email': email,
          'password': password,
        }),
      );

      if (response.statusCode == 200) {
        final responseData = jsonDecode(response.body);
        final token = responseData["token"];

        await _storage.write(key: 'auth_token', value: token);
        notifyListeners();
        return true;
      } else {
        print(response.body);
        return false;
      }
    } catch (e) {
      print('Erreur lors de la connexion: $e');
      rethrow;
    }
  }

  Future<bool> logout() async {
    try {
      final token = await _storage.read(key: 'auth_token');

      if (token != null) {
        final response = await http.post(
          Uri.parse('$_baseUrl/logout'),
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
        );

        if (response.statusCode == 200) {
          await _storage.delete(key: 'auth_token');
          notifyListeners();
          return true;
        } else {
          print('Échec de la déconnexion: ${response.body}');
          return false;
        }
      } else {
        print('Aucun token trouvé');
        return false;
      }
    } catch (e) {
      print('Erreur lors de la déconnexion: $e');
      rethrow;
    }
  }

  Future<Map<String, dynamic>> Profile() async {
    try {
      final token = await _storage.read(key: 'auth_token');

      if (token != null) {
        final response = await http.get(
          Uri.parse('$_baseUrl/profile'),
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
        );

        if (response.statusCode == 200) {
          final profileData = jsonDecode(response.body);
          return profileData;
        } else {
          throw Exception('Failed to load profile');
        }
      } else {
        throw Exception('User not logged in');
      }
    } catch (e) {
      print('Erreur lors de la récupération du profil: $e');
      rethrow;
    }
  }
}
