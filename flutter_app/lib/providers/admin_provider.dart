import 'dart:convert';
import 'package:flutter/cupertino.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;

class AdminProvider with ChangeNotifier {
  final FlutterSecureStorage _storage = FlutterSecureStorage();
  final String _baseUrl = "http://localhost:8080";

  bool get isAuthenticated => _storage.read(key: 'auth_token') != null;

  bool get isAdmin => _storage.read(key: 'user_role') == 'admin'; 

  Future<List<User>> fetchUsers() async {
    
    try {
      final token = await _storage.read(key: 'auth_token');

      if (token != null && isAdmin) {
        final response = await http.get(
          Uri.parse('$_baseUrl/users'),
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
        );

        if (response.statusCode == 200) {
          final List<dynamic> usersJson = jsonDecode(response.body);
          List<User> users = usersJson.map((json) => User.fromJson(json)).toList();
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

      if (token != null && isAdmin) {
        final response = await http.put(
          Uri.parse('$_baseUrl/users/${user.id}'),
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

    if (token != null && isAdmin) {
      final response = await http.delete(
        Uri.parse('$_baseUrl/users/$userId'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      if (response.statusCode == 204) {
        return true; // Suppression réussie
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

class User {
  final int id;
  final String username;
  final String email;

  User({required this.id, required this.username, required this.email});

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'],
      username: json['username'],
      email: json['email'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'username': username,
      'email': email,
    };
  }
}
