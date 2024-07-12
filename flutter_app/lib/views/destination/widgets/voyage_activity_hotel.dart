import 'package:flutter/cupertino.dart';
import 'package:flutter_app/models/activity_model.dart';
import 'package:flutter_app/models/hotel_model.dart';
import 'package:flutter_app/views/destination/widgets/activity_list.dart';
import 'package:flutter_app/views/destination/widgets/hotel_list.dart';

class VoyageActivityHotelList extends StatelessWidget {
  final List<Activity> activities;
  final List<Activity> selectedActivities;
  final List<Hotel> hotels;
  final List<Hotel> selectedHotels;
  final Function toggleActivity;
  final Function toggleHotel;

  const VoyageActivityHotelList({
    Key? key,
    required this.activities,
    required this.selectedActivities,
    required this.hotels,
    required this.selectedHotels,
    required this.toggleActivity,
    required this.toggleHotel,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          ActivityList(
            activities: activities,
            selectedActivities: selectedActivities,
            toggleActivity: toggleActivity,
          ),
          SizedBox(height: 20),
          // Liste des hôtels
          HotelList(
            hotels: hotels,
            selectedHotels: selectedHotels,
            toggleHotel: toggleHotel,
          ),
        ],
      ),
    );
  }
}