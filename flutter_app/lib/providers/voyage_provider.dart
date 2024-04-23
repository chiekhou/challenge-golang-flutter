import 'dart:collection';
import 'dart:convert';
import 'dart:io';
import 'package:flutter/cupertino.dart';
import '../models/activity_model.dart';
import '../models/voyage_model.dart';
import 'package:http/http.dart' as http;

class VoyageProvider extends ChangeNotifier {
  final String host = '10.0.2.2'; // version emulateur
  // final String host = 'localhost'; // version web
  List<Voyage> _voyages = [];
  bool isLoading = false;

  UnmodifiableListView<Voyage> get voyages => UnmodifiableListView(_voyages);

  Future<void> fetchData() async {
    try {
      isLoading = true;
      http.Response response = await http.get(Uri.http(host, '/api/voyages'));
      if (response.statusCode == 200) {
        _voyages = (json.decode(response.body) as List)
            .map((voyageJson) => Voyage.fromJson(voyageJson))
            .toList();
        isLoading = false;
        notifyListeners();
      }
    } catch (e) {
      isLoading = false;
      rethrow;
    }
  }

  Future<void> addVoyage(Voyage voyage) async {
    try {
      http.Response response = await http.post(
        Uri.http(host, '/api/voyage'),
        body: json.encode(
          voyage.toJson(),
        ),
        headers: {'Content-type': 'application/json'},
      );
      if (response.statusCode == 200) {
        _voyages.add(
          Voyage.fromJson(
            json.decode(response.body),
          ),
        );
        notifyListeners();
      }
    } catch (e) {
      rethrow;
    }
  }

  Future<void> updateVoyage(Voyage voyage, String activityId) async {
    try {
      Activity activity =
          voyage.activities.firstWhere((activity) => activity.id == activityId);
      activity.status = ActivityStatus.done;
      http.Response response = await http.put(
        Uri.http(host, '/api/voyage'),
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

  Voyage getById(String id) {
    return voyages.firstWhere((voyage) => voyage.id == id);
  }
}
