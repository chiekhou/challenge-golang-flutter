import 'dart:collection';
import 'dart:convert';
import 'package:flutter/cupertino.dart';
import 'package:flutter_app/config/app_config.dart';
import 'package:http_parser/http_parser.dart';
import '../models/activity_model.dart';
import '../models/destination_model.dart';
import 'package:http/http.dart' as http;
import 'package:path/path.dart';
import 'dart:io';

class DestinationProvider extends ChangeNotifier {
  List<Destination> _destinations = [];
  bool isLoading = false;
  final apiAuthority = AppConfig.getApiAuthority();
  final isSecure = AppConfig.isSecure();

  UnmodifiableListView<Destination> get destinations =>
      UnmodifiableListView(_destinations);

  Destination getDestinationByName(String destinationName) => destinations
      .firstWhere((destination) => destination.name == destinationName);

  UnmodifiableListView<Destination> getFilteredDestinations(String filter) =>
      UnmodifiableListView(
        _destinations
            .where(
              (destination) => destination.name.toLowerCase().startsWith(
                    filter.toLowerCase(),
                  ),
            )
            .toList(),
      );

  Future<void> fetchData() async {
    try {
      isLoading = true;
      notifyListeners();

      final url = isSecure
          ? Uri.https(apiAuthority, '/api/destinations')
          : Uri.http(apiAuthority, '/api/destinations');
      final response = await http.get(url);
      print('Fetching data from $url');
      print('Response body: ${response.body}');

      if (response.statusCode == 200) {
        final responseData = json.decode(response.body);
        if (responseData is List) {
          _destinations = responseData
              .map((destinationJson) => Destination.fromJson(destinationJson))
              .toList();
        } else if (responseData is Map && responseData.containsKey('data')) {
          _destinations = (responseData['data'] as List)
              .map((voyageJson) => Destination.fromJson(voyageJson))
              .toList();
        } else {
          throw Exception('Unexpected response format');
        }
      } else {
        throw Exception('Failed to load destinations');
      }
    } catch (e) {
      print('Error: $e ');
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }

  Future<void> addActivityToDestination(Activity newActivity) async {
    try {
      int destinationId = getDestinationByName(newActivity.destination).id;
      final url = isSecure
          ? Uri.https(apiAuthority, '/api/destination/$destinationId/activity')
          : Uri.http(apiAuthority, '/api/destination/$destinationId/activity');
      http.Response response = await http.post(url,
        headers: {'Content-type': 'application/json'},
        body: json.encode(
          newActivity.toJson(),
        ),
      );
      if (response.statusCode == 200) {
        int index = _destinations
            .indexWhere((destination) => destination.id == destinationId);
        _destinations[index] = Destination.fromJson(
          json.decode(response.body),
        );
        notifyListeners();
      }
    } catch (e) {
      rethrow;
    }
  }

  Future<dynamic> verifyIfActivityNameIsUnique(
      String destinationName, String activityName) async {
    try {
      Destination destination = getDestinationByName(destinationName);
      final url = isSecure
          ? Uri.https(apiAuthority, '/api/destination/${destination.id}/activities/verify/$activityName')
          : Uri.http(apiAuthority,'/api/destination/${destination.id}/activities/verify/$activityName');
      http.Response response = await http.get(url);
      if (response.statusCode != 200) {
        return json.decode(response.body);
      } else {
        return null;
      }
    } catch (e) {
      rethrow;
    }
  }

  Future<String> uploadImage(File pickedImage) async {
    try {
      final url = isSecure
          ? Uri.https(apiAuthority, '/api/activity/image')
          : Uri.http(apiAuthority,'/api/activity/image');
      var request =
          http.MultipartRequest("POST", url);
      request.files.add(
        http.MultipartFile.fromBytes(
          'activity',
          pickedImage.readAsBytesSync(),
          filename: basename(pickedImage.path),
          contentType: MediaType("multipart", "form-data"),
        ),
      );
      var response = await request.send();
      if (response.statusCode == 200) {
        var responseData = await response.stream.bytesToString();
        return json.decode(responseData);
      } else {
        throw 'error';
      }
    } catch (e) {
      rethrow;
    }
  }
}
