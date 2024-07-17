import 'package:flutter_app/models/hotel_model.dart';
import 'package:intl/intl.dart';

import 'activity_model.dart';

class Voyage {
  int? id;
  String destination;
  List<Activity> activities;
  List<Hotel> hotels;
  DateTime? dateAller;
  DateTime? dateRetour;
  Voyage({
    required this.destination,
    required this.activities,
    required this.hotels,
    this.dateAller,
    this.dateRetour,
    this.id,
  });

  @override
  String toString() {
    return 'Voyage{id: $id, destination: $destination, dateAller: $dateAller,dateRetour: $dateRetour, activities: $activities, hotels: $hotels}';
  }

  factory Voyage.fromJson(Map<String, dynamic> json) {
    var activitiesFromJson = json['activities'] as List;
    List<Activity> activityList = activitiesFromJson
        .map((activityJson) =>
            Activity.fromJson(activityJson as Map<String, dynamic>))
        .toList();

    var hotelsFromJson = json['hotels'] as List;
    List<Hotel> hotelList = hotelsFromJson
        .map((hotelJson) => Hotel.fromJson(hotelJson as Map<String, dynamic>))
        .toList();

    return Voyage(
      id: json['id'],
      activities: activityList,
      hotels: hotelList,
      dateAller: DateTime.parse(json['dateAller']),
      dateRetour: DateTime.parse(json['dateRetour']),
      destination: json['destination'],
    );
  }
  Map<String, dynamic> toJson() {
    final DateFormat formatter = DateFormat("yyyy-MM-ddTHH:mm:ss'Z'");
    final String? dateFormattedA =
        dateAller != null ? formatter.format(dateAller!) : null;
    final String? dateFormattedR =
        dateRetour != null ? formatter.format(dateRetour!) : null;
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
