import 'dart:convert';
import 'package:flutter/cupertino.dart';
import 'package:flutter_app/config/app_config.dart';
import 'package:flutter_app/models/user_model.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;

class AdminProvider with ChangeNotifier {
  final FlutterSecureStorage _storage = FlutterSecureStorage();
  final apiAuthority = AppConfig.getApiAuthority();
  final isSecure = AppConfig.isSecure();

  Future<List<User>> fetchUsers() async {
    try {
      final token = await _storage.read(key: 'auth_token');

      final url = isSecure
          ? Uri.https(apiAuthority, '/api/users')
          : Uri.http(apiAuthority, '/api/users');

      if (token != null) {
        final response = await http.get(
          url,
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
        );

        if (response.statusCode == 200) {
          final List<dynamic> usersJson = jsonDecode(response.body);
          List<User> users =
              usersJson.map((json) => User.fromJson(json)).toList();
          return users;
        } else {
          throw Exception('Failed to fetch users');
        }
      } else {
        throw Exception('Unauthorized to fetch users');
      }
    } catch (e) {
      print('Error fetching users: $e');
      rethrow;
    }
  }

  Future<bool> updateUser(User user) async {
    try {
      final token = await _storage.read(key: 'auth_token');

      final url = isSecure
          ? Uri.https(apiAuthority, '/api/users/${user.id}')
          : Uri.http(apiAuthority, '/api/users/${user.id}');

      if (token != null) {
        final response = await http.put(
          url,
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
          body: jsonEncode(user.toJson()),
        );

        if (response.statusCode == 200) {
          return true;
        } else {
          print('Failed to update user: ${response.statusCode}');
          return false;
        }
      } else {
        throw Exception('Unauthorized to update user');
      }
    } catch (e) {
      print('Error updating user: $e');
      rethrow;
    }
  }

  Future<bool> deleteUser(int userId) async {
    try {
      final token = await _storage.read(key: 'auth_token');

      final url = isSecure
          ? Uri.https(apiAuthority, '/api/users/$userId')
          : Uri.http(apiAuthority, '/api/users/$userId');

      if (token != null) {
        final response = await http.delete(
          url,
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
        );

        if (response.statusCode == 204) {
          return true;
        } else {
          print('Failed to delete user: ${response.statusCode}');
          return false;
        }
      } else {
        throw Exception('Unauthorized to delete user');
      }
    } catch (e) {
      print('Error deleting user: $e');
      rethrow;
    }
  }
}
