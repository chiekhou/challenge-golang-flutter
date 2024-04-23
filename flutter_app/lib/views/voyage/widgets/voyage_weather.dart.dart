import 'package:flutter/material.dart';
import 'package:flutter_app/widgets/app_loader.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class VoyageWeather extends StatelessWidget {
  final String destinationName;
  final String hostBase = 'https://api.openweathermap.org/data/2.5/weather?q=';
  final String apiKey = '&appid=f5dd5d8df05953a6da3b3676bf708ee0';

  const VoyageWeather({super.key, required this.destinationName});

  String get query => '$hostBase$destinationName$apiKey';

  Future<String> get getWeather {
    return http.get(Uri.parse(query)).then((http.Response response) {
      Map<String, dynamic> body = json.decode(response.body);
      return body['weather'][0]['icon'] as String;
    }).catchError((e) => 'error');
  }

  String getIconUrl(String iconName) {
    return 'https://openweathermap.org/img/wn/$iconName@2x.png';
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: getWeather,
      builder: (_, snapshot) {
        if (snapshot.hasError) {
          return const Text('error');
        } else if (snapshot.hasData) {
          return Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 5),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: <Widget>[
                const Text(
                  'Météo',
                  style: TextStyle(fontSize: 20),
                ),
                Image.network(
                  getIconUrl(snapshot.data as String),
                  width: 50,
                  height: 50,
                ),
              ],
            ),
          );
        } else {
          return const AppLoader();
        }
      },
    );
  }
}
