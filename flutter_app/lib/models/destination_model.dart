import 'package:flutter_app/models/hotel_model.dart';

import 'activity_model.dart';

class Destination {
  int id;
  String image;
  String name;
  List<Activity> activities;
  List<Hotel> hotels;
  Destination({
    required this.id,
    required this.image,
    required this.name,
    required this.activities,
    required this.hotels,
  });

  Destination.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        image = json['image'],
        name = json['name'],
        activities = (json['activities'] as List)
            .map((activityJson) => Activity.fromJson(activityJson))
            .toList(),
        hotels = (json['hotels'] as List)
            .map((hotelJson) => Hotel.fromJson(hotelJson))
            .toList();
}
