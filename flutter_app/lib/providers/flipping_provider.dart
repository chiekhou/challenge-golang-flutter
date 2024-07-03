import 'dart:convert';
import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;

class FlippingToggle extends ChangeNotifier {
  final String host = '10.0.2.2:8080'; // version emulateur
  // final String host = 'localhost:8080'; // version web
  final bool enabled;

  FlippingToggle({required this.enabled});

  factory FlippingToggle.fromJson(Map<String, dynamic> json) {
    return FlippingToggle(
      enabled: json['enabled'],
    );
  }
}

Future<FlippingToggle> fetchFeatureToggles() async {
  final response = await http.get(Uri.parse('http://10.0.2.2:8080/api/flipping/feature?feature=active_voyage'));

  if (response.statusCode == 200) {
    try {
      return FlippingToggle.fromJson(jsonDecode(response.body));
    } catch (e) {
      print('Erreur lors de la conversion JSON: $e');
      throw Exception('Erreur lors de la conversion JSON');
    }
  } else {
    print('Erreur lors de la récupération des basculements de fonctionnalité: ${response.body}');
    throw Exception('Erreur lors de la récupération des basculements de fonctionnalité');
  }
}

Future<void> updateFeatureToggle( bool enabled) async {
  final response = await http.put(Uri.parse('http://10.0.2.2:8080/api/flipping/feature?feature=active_voyage'),
    headers: {
      'Content-Type': 'application/json',
    },
    body: json.encode({
      'enabled': enabled,
    }),
  );

  if (response.statusCode != 200) {
    throw Exception('Failed to update feature toggle');
  }
}
