import 'dart:collection';
import 'dart:convert';
import 'package:flutter/cupertino.dart';
import 'package:flutter_app/config/app_config.dart';
import 'package:flutter_app/models/groupe_model.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;

class GroupVoyageProvider extends ChangeNotifier {

  final String host = "10.0.2.2";
  final FlutterSecureStorage _storage = FlutterSecureStorage();
  List<Groupe> _groupes = [];
  UnmodifiableListView<Groupe> get groupes => UnmodifiableListView(_groupes)
  final FlutterSecureStorage _storage = const FlutterSecureStorage();
  final apiAuthority = AppConfig.getApiAuthority();
  final isSecure = AppConfig.isSecure();
  List<Groupe> _groupes = [];
  UnmodifiableListView<Groupe> get groupes => UnmodifiableListView(_groupes);

  
  Future<List<Groupe>> fetchGroupes() async {
    try {
      final token = await _storage.read(key: 'auth_token');
      final url = isSecure
          ? Uri.https(apiAuthority, '/groupes')
          : Uri.http(apiAuthority, '/groupes');

      if (token != null) {
        final response = await http.get(
          url,
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
        );

        if (response.statusCode == 200) {
          final Map<String, dynamic> responseBody = jsonDecode(response.body);

          if (responseBody.containsKey('data') &&
              responseBody['data'] is List) {
            final List<dynamic> groupsJson = responseBody['data'];
            List<Groupe> groups =
                groupsJson.map((json) => Groupe.fromJson(json)).toList();
            return groups;
          } else {
            return []; // Retourne une liste vide si 'data' est absent ou null
          }
        } else {
          throw Exception('Failed to fetch groups');
        }
      } else {
        throw Exception('Unauthorized to fetch groups');
      }
    } catch (e) {
      print('Error fetching groups: $e');
      rethrow;
    }
  }
        
        
  Future<bool> CreateGroup(double budget, String nom, int? voyageId) async {
    try {
      final url = isSecure
          ? Uri.https(apiAuthority, '/create_group')
          : Uri.http(apiAuthority, '/create_group');
      String? token = await _storage.read(key: 'auth_token');
      if (token != null) {
        final response = await http.post(url,
            headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer $token'
            },
          body: jsonEncode({
            'budget': budget,
            'nom': nom,
            'voyage_id': voyageId
          }),
        );
        return response.statusCode == 200;
      } else {
        throw Exception('User not logged in');
      }
    } catch (e) {
      rethrow;
    }
  }


  //Update le budget
  Future<bool> UpdateBudget(int ID, double budget) async {
    try {
      final url = isSecure
          ? Uri.https(apiAuthority, '/groupes/$ID/update_budget')
          : Uri.http(apiAuthority, '/groupes/$ID/update_budget');

      String? token = await _storage.read(key: 'auth_token');

      if (token != null) {
        final response = await http.put(url,
            headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer $token'
            },
            body: jsonEncode({"budget": budget}));

        if (response.statusCode == 200) {
          return true;
        } else {
          return false;
        }
      } else {
        throw Exception('user not logged in');
      }
    } catch (e) {
      rethrow;
    }
  }

  // Récupérer les groupes par user
  Future<void> GetGroups() async {

    try {
       
    String? token = await _storage.read(key: 'auth_token');
    if (token != null) {
        final url = isSecure
        ? Uri.https(apiAuthority, '/groupes/my_groups')
        : Uri.http(apiAuthority, '/groupes/my_groups');
      final response = await http.get(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token'
        },
      );
        if (response.statusCode == 200) {
          final responseData = jsonDecode(response.body);
          if (responseData is List) {
            _groupes = responseData.map((groupeJson) => Groupe.fromJson(groupeJson)).toList();
          } else if (responseData is Map && responseData.containsKey('data')) {
            _groupes = (responseData['data'] as List).map((groupJson) => Groupe.fromJson(groupJson)).toList();
          } else {
            throw Exception('Unexpected response format');
          }
        } else {
          throw Exception('No groups found');
        }
      } else {
        throw Exception('User not logged in');
      }
    } catch (e) {
      rethrow;
    } finally {
      notifyListeners();
    }
  }

  Groupe getGroupeById(int id) {
    return _groupes.firstWhere((groupe) => groupe.id == id);
  }


  //Groupe par ID
  Future<Map<String, dynamic>> GetGroupByID(int ID) async {
    final url = isSecure
        ? Uri.https(apiAuthority, '/groupes/$ID')
        : Uri.http(apiAuthority, '/groupes/$ID');
    String? token = await _storage.read(key: 'auth_token');
    if (token != null) {
      final response = await http.get(url, headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token'
      });

      if (response.statusCode == 200) {
        return jsonDecode(response.body);
      } else {
        throw Exception('Group not found');
      }
    } else {
      throw Exception('User not logged in');
    }
  }

        
  Future<bool> SendInvitation(int ID, String email) async {
    try {
      String? token = await _storage.read(key: 'auth_token');

      final url = isSecure
          ? Uri.https(apiAuthority, '/groupes/$ID/send_invitation')
          : Uri.http(apiAuthority, '/groupes/$ID/send_invitation');

      if (token != null) {
        final response = await http.post(url,
            headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer $token'
            },
            body: jsonEncode({"email": email}));
        if (response.statusCode == 200) {
          return true;
        } else {
          return false;
        }
      } else {
        throw Exception('User not logged in');
      }
    } catch (e) {
      rethrow;
    }
  }

        
        
  Future<bool> deleteGroup(int groupeId) async {
    try {
      final token = await _storage.read(key: 'auth_token');
      final url = isSecure
          ? Uri.https(apiAuthority, '/groupes/${groupeId}/delete_group')
          : Uri.http(apiAuthority, '/groupes/${groupeId}/delete_group');

      if (token != null) {
        final response = await http.delete(
          url,
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
        );
        if (response.statusCode == 200 || response.statusCode == 204) {
          return true;
        } else {
          print('Failed to delete groupe: ${response.statusCode}');
          return false;
        }
      } else {
        throw Exception('Unauthorized to delete groupe');
      }
    } catch (e) {
      print('Error deleting groupe: $e');

      rethrow;
    }
  }
}
