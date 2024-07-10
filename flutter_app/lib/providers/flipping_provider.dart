import 'dart:convert';
import 'package:flutter/cupertino.dart';
import 'package:flutter_app/config/app_config.dart';
import 'package:http/http.dart' as http;

class FlippingToggle extends ChangeNotifier {

  final bool enabled;

  FlippingToggle({required this.enabled});

  factory FlippingToggle.fromJson(Map<String, dynamic> json) {
    return FlippingToggle(
      enabled: json['enabled'],
    );
  }
}

final apiAuthority = AppConfig.getApiAuthority();
final isSecure = AppConfig.isSecure();
final url = isSecure
    ? Uri.https(apiAuthority, 'api/flipping/feature?feature=active_voyage')
    : Uri.http(apiAuthority, 'api/flipping/feature?feature=active_voyage');

Future<FlippingToggle> fetchFeatureToggles() async {

  final response = await http.get(url);

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
  final response = await http.put(url,
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
