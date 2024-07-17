import 'dart:collection';
import 'dart:convert';

import 'dart:io';
import 'package:flutter/cupertino.dart';
import 'package:flutter_app/models/hotel_model.dart';
import '../models/activity_model.dart';
import '../models/voyage_model.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;

class VoyageProvider extends ChangeNotifier {
  final String host = '10.0.2.2:8080'; // version emulateur
  //final String host = 'localhost:8080'; // version web
  List<Voyage> _voyages = [];
  bool isLoading = false;
  final FlutterSecureStorage _storage = FlutterSecureStorage();

  UnmodifiableListView<Voyage> get voyages => UnmodifiableListView(_voyages);
  Future<void> fetchData() async {
    try {
      String? token = await _storage.read(key: 'auth_token');
      if(token != null){
        isLoading = true;
        notifyListeners();

        final response = await http.get(
          Uri.http(host,'/api/voyages'),
            headers:{
              'Content-Type': 'application/json',
              'Authorization': 'Bearer $token'
            }
        );
        // print('Fetching data from $url');
        // print('Response body: ${response.body}');

        if (response.statusCode == 200) {
          final responseData = json.decode(response.body);
          if (responseData is List) {
            _voyages = responseData
                .map((voyageJson) => Voyage.fromJson(voyageJson))
                .toList();
          } else if (responseData is Map && responseData.containsKey('data')) {
            _voyages = (responseData['data'] as List)
                .map((voyageJson) => Voyage.fromJson(voyageJson))
                .toList();
          } else {
            throw Exception('Unexpected response format');
          }
        } else {
          throw Exception('Failed to load voyages');
        }
      }else{
        print('Aucun token trouvé');
        throw Exception('No token found');
      }
    } catch (e) {
      // print('Error: $e');
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }

  Future<void> addVoyage(Voyage voyage) async {
    try {
      String? token = await _storage.read(key: 'auth_token');
      if(token != null){
        final jsonData = json.encode(voyage.toJson());
        print('Données envoyées : $jsonData');

        final response = await http.post(
          Uri.http(host, '/api/voyages'),
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token'
          },
          body: jsonData,
        );

        if (response.statusCode == 200) {
          final responseData = json.decode(response.body);
          // print('Réponse reçue : $responseData');

          final Voyage newVoyage = Voyage.fromJson(responseData['data']);
          // print('Nouveau voyage créé : ${newVoyage.id}');

          _voyages.add(newVoyage);
          notifyListeners();
        } else {
          print('Erreur lors de l\'ajout du voyage : ${response.body}');
          throw Exception(
              'Erreur lors de l\'ajout du voyage : ${response.statusCode}');
        }
      }else{
        print('Aucun token trouvé');
        throw Exception('No token found');
      }
    } catch (e) {
      throw Exception('Erreur lors de l\'ajout du voyage : $e');
    }
  }

  Future<void> updateVoyage(Voyage voyage, int activityId) async {
    try {
      String? token = await _storage.read(key: 'auth_token');
      if(token != null){
        Activity activity =
        voyage.activities.firstWhere((activity) => activity.id == activityId);
        activity.status = ActivityStatus.done;
        http.Response response = await http.put(
          Uri.http(host, '/api/voyages'),
          body: json.encode(
            voyage.toJson(),
          ),
          headers: {
            'Content-type': 'application/json',
            'Authorization': 'Bearer $token'
          },
        );
        if (response.statusCode != 200) {
          activity.status = ActivityStatus.ongoing;
          throw const HttpException('error');
        }
      }else{
        print('Aucun token trouvé');
        throw Exception('No token found');
      }
      notifyListeners();
    } catch (e) {
      rethrow;
    }
  }


  Future<void> updateVoyageHotel(Voyage voyage, int hotelId) async {
    try {
      String? token = await _storage.read(key: 'auth_token');
      if(token != null){
        Hotel hotel =
        voyage.hotels.firstWhere((hotel) => hotel.id == hotelId);
        hotel.status = HotelStatus.done;
        http.Response response = await http.put(
          Uri.http(host, '/api/voyages'),
          body: json.encode(
            voyage.toJson(),
          ),
          headers: {
            'Content-type': 'application/json',
            'Authorization': 'Bearer $token'
          },
        );
        if (response.statusCode != 200) {
          hotel.status = HotelStatus.ongoing;
          throw const HttpException('error');
        }
        notifyListeners();
      }else{
        print('Aucun token trouvé');
        throw Exception('No token found');
      }
    } catch (e) {
      rethrow;
    }
  }

  Voyage getById(int id) {
    return voyages.firstWhere((voyage) => voyage.id == id);
  }

  Activity getActivityByIds({required int activityId, required int voyageId}) {
    return getById(voyageId)
        .activities
        .firstWhere((activity) => activity.id == activityId);
  }


  Hotel getHotelByIds({required int hotelId, required int voyageId}) {
    return getById(voyageId)
        .hotels
        .firstWhere((hotel) => hotel.id == hotelId);
  }
}
