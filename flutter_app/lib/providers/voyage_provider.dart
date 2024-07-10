import 'dart:collection';
import 'dart:convert';

import 'dart:io';
import 'package:flutter/cupertino.dart';
import 'package:flutter_app/config/app_config.dart';
import 'package:flutter_app/models/hotel_model.dart';
import '../models/activity_model.dart';
import '../models/voyage_model.dart';
import 'package:http/http.dart' as http;

class VoyageProvider extends ChangeNotifier {

  List<Voyage> _voyages = [];
  bool isLoading = false;
  final apiAuthority = AppConfig.getApiAuthority();
  final isSecure = AppConfig.isSecure();


  UnmodifiableListView<Voyage> get voyages => UnmodifiableListView(_voyages);
  Future<void> fetchData() async {

    try {
      isLoading = true;
      notifyListeners();

      final url = isSecure
          ? Uri.https(apiAuthority, '/api/voyages')
          : Uri.http(apiAuthority, '/api/voyages');

      final response = await http.get(url);
       print('Fetching data from $url');
      print('Response body: ${response.body}');

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
    } catch (e) {
      print('Error: $e');
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }

  Future<void> addVoyage(Voyage voyage) async {
    try {
      final jsonData = json.encode(voyage.toJson());
      print('Données envoyées : $jsonData');

      final url = isSecure
          ? Uri.https(apiAuthority, '/api/voyages')
          : Uri.http(apiAuthority, '/api/voyages');
      final response = await http.post(url,
        headers: {'Content-Type': 'application/json'},
        body: jsonData,
      );

      if (response.statusCode == 200) {
        final responseData = json.decode(response.body);
        ('Réponse reçue : $responseData');

        final Voyage newVoyage = Voyage.fromJson(responseData['data']);
        print('Nouveau voyage créé : ${newVoyage.id}');

        _voyages.add(newVoyage);
        notifyListeners();
      } else {
        print('Erreur lors de l\'ajout du voyage : ${response.body}');
        throw Exception(
            'Erreur lors de l\'ajout du voyage : ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Erreur lors de l\'ajout du voyage : $e');
    }
  }

  Future<void> updateVoyage(Voyage voyage, int activityId) async {
    try {
      Activity activity =
          voyage.activities.firstWhere((activity) => activity.id == activityId);
      activity.status = ActivityStatus.done;
      final url = isSecure
          ? Uri.https(apiAuthority, '/api/voyages')
          : Uri.http(apiAuthority, '/api/voyages');
      http.Response response = await http.put(url,
        body: json.encode(
          voyage.toJson(),
        ),
        headers: {'Content-type': 'application/json'},
      );
      if (response.statusCode != 200) {
        activity.status = ActivityStatus.ongoing;
        throw const HttpException('error');
      }
      notifyListeners();
    } catch (e) {
      rethrow;
    }
  }


  Future<void> updateVoyageHotel(Voyage voyage, int hotelId) async {
    try {
      Hotel hotel =
      voyage.hotels.firstWhere((hotel) => hotel.id == hotelId);
      hotel.status = HotelStatus.done;
      final url = isSecure
          ? Uri.https(apiAuthority, '/api/voyages/hotel')
          : Uri.http(apiAuthority, '/api/voyages/hotel');
      http.Response response = await http.put(url,
        body: json.encode(
          voyage.toJson(),
        ),
        headers: {'Content-type': 'application/json'},
      );
      if (response.statusCode != 200) {
        hotel.status = HotelStatus.ongoing;
        throw const HttpException('error');
      }
      notifyListeners();
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
