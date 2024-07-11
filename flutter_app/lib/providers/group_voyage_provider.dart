import 'dart:collection';
import 'dart:convert';

import 'package:flutter/cupertino.dart';
import 'package:flutter_app/models/groupe_model.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;


class GroupVoyageProvider extends ChangeNotifier{
  final String host = "10.0.2.2";
  final FlutterSecureStorage _storage = FlutterSecureStorage();
  final String _baseUrl = "http://localhost:8080";
  List<Groupe> _groupes = [];
  UnmodifiableListView<Groupe> get groupes => UnmodifiableListView(_groupes);


  //Créer un groupe
  Future<bool>CreateGroup(double budget, String nom)async{
    try{
      String? token = await _storage.read(key: 'auth_token');
      if(token != null){

        final response = await http.post(
            Uri.parse('http://$host:8080/create_group'),
            headers: {
              'Content-Type' : 'application/json',
              'Authorization' : 'Bearer $token'
            },
            body: jsonEncode({
              'budget': budget,
              'nom': nom
            })
        );
        if(response.statusCode == 200){
          return true;
        }else{
          return false;
        }
      }else{
        throw Exception('User not logged in');
      }
      }catch(e){
      rethrow;
    }
  }

  //Update le budget
  Future<bool>UpdateBudget(int ID, double budget)async{
    try{
      String? token = await _storage.read(key: 'auth_token');

      if(token != null){
        final response = await http.put(
          Uri.parse('http://$host:8080/groupes/$ID/update_budget'),
          headers:{
            'Content-Type' : 'application/json',
            'Authorization' : 'Bearer $token'
          },
          body: jsonEncode({
            "budget" : budget
          })
        );

        if(response.statusCode == 200){
          return true;
        }else{
          return false;
        }
      }else{
        throw Exception('user not logged in');
      }
    }catch(e){
      rethrow;
    }
  }

  //Récupérer les groupes par user
  Future<void> GetGroups() async {
    String? token = await _storage.read(key: 'auth_token');
    if (token != null) {
      final response = await http.get(
        Uri.parse('http://$host:8080/groupes/my_groups'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token'
        },
      );
      if (response.statusCode == 200) {
        final responseData = jsonDecode(response.body);
        if (responseData is List) {
          _groupes = responseData.map((groupeJson) => Groupe.fromJson(groupeJson)).toList();
        } else {
          throw Exception('Unexpected response format');
        }
        notifyListeners();
      } else {
        throw Exception('No groups found');
      }
    } else {
      throw Exception('User not logged in');
    }
  }

  Groupe getGroupeById(int id) {
    return _groupes.firstWhere((groupe) => groupe.id == id);
  }

  //Groupe par ID
  Future<Map<String, dynamic>>GetGroupByID(int ID)async{
    String? token = await _storage.read(key: 'auth_token');
    if(token != null){
      final response = await http.get(
        Uri.parse('http://$host:8080/groupes/$ID'),
        headers: {
          'Content-Type' : 'application/json',
          'Authorization' : 'Bearer $token'
        }
      );
      if(response.statusCode == 200){
        return jsonDecode(response.body);
      }else{
        throw Exception('Group not found');
      }
    }else{
      throw Exception('User not logged in');
    }
  }

  //Inviter un user
  Future<bool>SendInvitation(int ID,String email)async{
    try{
      String? token = await _storage.read(key: 'auth_token');
      if(token != null){
        final response = await http.post(
            Uri.parse('http://$host:8080/groupes/$ID/send_invitation'),
            headers: {
              'Content-Type' : 'application/json',
              'Autorization' : 'Bearer $token'
            },
            body: jsonEncode({
              "email" : email
            })
        );
        if(response.statusCode == 201){
          return true;
        }else{
          return false;
        }
      }else{
        throw Exception('User not logged in');
      }
    }catch(e){
      rethrow;
    }
  }

  Future<void>JoinGroup(int _gorupId, String? _token)async{
    try{
      final response = await http.get(
        Uri.parse('http://$host:8080/groupes/$_gorupId/join?token=$_token'),
        headers: {
          'Content-Type': 'application/json',
          //'Authorization': 'Bearer $authToken',
        },
      );
      if (response.statusCode != 200) {
        throw Exception('Erreur lors de la tentative de rejoindre le groupe');
      }
      notifyListeners();
    }catch(e){
      rethrow;
    }
  }
}