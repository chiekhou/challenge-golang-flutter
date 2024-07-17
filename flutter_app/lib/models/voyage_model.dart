import 'package:flutter_app/models/hotel_model.dart';
import 'package:intl/intl.dart';

import 'activity_model.dart';

class Voyage {
  final int? id;
  String destination;
  List<Activity> activities;
  List<Hotel> hotels;
  DateTime? dateAller;
  DateTime? dateRetour;
  final int? userId;

  Voyage({
    this.id,
    required this.destination,
    required this.activities,
    required this.hotels,
    this.dateAller,
    this.dateRetour,
    this.userId,
  });


  @override
  String toString() {
    return 'Voyage{id: $id, destination: $destination, dateAller: $dateAller,dateRetour: $dateRetour, activities: $activities, hotels: $hotels}';
  }

  factory Voyage.fromJson(Map<String, dynamic> json) {
    return Voyage(
      id: json['id'],
      destination: json['destination'] ?? '',
      activities: json['activities'] != null
          ? (json['activities'] as List).map((i) => Activity.fromJson(i)).toList()
          : [],
      hotels: json['hotels'] != null
          ? (json['hotels'] as List).map((i) => Hotel.fromJson(i)).toList()
          : [],
      dateAller: DateTime.parse(json['dateAller']),
      dateRetour: DateTime.parse(json['dateRetour']),
      userId: json['user_id'],
    );
  }
  Map<String, dynamic> toJson() {
    final DateFormat formatter = DateFormat("yyyy-MM-ddTHH:mm:ss'Z'");
    final String? dateFormattedA = dateAller != null ? formatter.format(dateAller!) : null;
    final String? dateFormattedR = dateRetour != null ? formatter.format(dateRetour!) : null;
    return {
      'id': id,
      'activities': activities.map((activity) => activity.toJson()).toList(),
      'hotels': hotels.map((hotel) => hotel.toJson()).toList(),
      'dateAller': dateFormattedA,
      'dateRetour': dateFormattedR,
      'destination': destination,
    };
  }
}
