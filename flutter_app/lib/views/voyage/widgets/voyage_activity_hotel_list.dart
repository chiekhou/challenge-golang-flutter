import 'package:flutter/material.dart';
import 'package:flutter_app/models/hotel_model.dart';
import 'package:flutter_app/views/google_map/google_map_hotel_view.dart';
import 'package:provider/provider.dart';
import '../../../models/activity_model.dart';
import '../../../models/voyage_model.dart';
import '../../../providers/voyage_provider.dart';
import '../../google_map/google_map_view.dart';

class VoyageActivityHotelList extends StatelessWidget {
  final int voyageId;
  final ActivityStatus filter;
  final HotelStatus filterh;
  VoyageActivityHotelList(
      {super.key, required this.voyageId, required this.filter , required this.filterh});


  @override
  Widget build(BuildContext context) {
    return Consumer<VoyageProvider>(
      builder: (context, voyageProvider, child) {
        final Voyage voyage = Provider.of<VoyageProvider>(context).getById(voyageId);
        final List<Activity> activities = voyage.activities
            .where((activity) => activity.status == filter)
            .toList();
        final List<Hotel> hotels = voyage.hotels
            .where((hotel) => hotel.status == filterh)
            .toList();

        return SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              if (filter == ActivityStatus.done && activities.isEmpty && hotels.isEmpty)
                Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: Center(
                    child: Text(
                      'Aucun voyage n\'est encore terminé',
                      style: TextStyle(fontSize: 16, color: Colors.grey),
                    ),
                  ),
                ),
              if (activities.isNotEmpty) ...[
                Container(
                  padding: const EdgeInsets.all(8.0),
                  child: Text(
                    'Liste des Activités',
                    style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                  ),
                ),
                ListView.builder(
                  shrinkWrap: true,
                  physics: NeverScrollableScrollPhysics(),
                  itemCount: activities.length,
                  itemBuilder: (context, i) {
                    Activity activity = activities[i];
                    return Container(
                      margin: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
                      child: filter == ActivityStatus.ongoing
                          ? Dismissible(
                        direction: DismissDirection.endToStart,
                        background: Container(
                          padding: const EdgeInsets.symmetric(horizontal: 15),
                          alignment: Alignment.centerRight,
                          decoration: BoxDecoration(
                            borderRadius: BorderRadius.circular(10),
                            color: Colors.greenAccent[700],
                          ),
                          child: const Icon(
                            Icons.check,
                            color: Colors.white,
                            size: 30,
                          ),
                        ),
                        key: ValueKey(activity.id),
                        child: InkWell(
                          onTap: () => Navigator.pushNamed(
                            context,
                            GoogleMapView.routeName,
                            arguments: {
                              'activityId': activity.id,
                              'voyageId': voyage.id
                            },
                          ),
                          child: Card(
                            child: ListTile(
                              title: Text(activity.name),
                            ),
                          ),
                        ),
                        confirmDismiss: (_) {
                          return Provider.of<VoyageProvider>(context, listen: false)
                              .updateVoyage(voyage, activity.id)
                              .then((_) => true)
                              .catchError((_) => false);
                        },
                      )
                          : Card(
                        child: ListTile(
                          title: Text(
                            activity.name,
                            style: const TextStyle(color: Colors.grey),
                          ),
                        ),
                      ),
                    );
                  },
                ),
              ],
              if (hotels.isNotEmpty) ...[
                Container(
                  padding: const EdgeInsets.all(8.0),
                  child: Text(
                    'Liste des Hôtels',
                    style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                  ),
                ),
                ListView.builder(
                  shrinkWrap: true,
                  physics: NeverScrollableScrollPhysics(),
                  itemCount: hotels.length,
                  itemBuilder: (context, i) {
                    Hotel hotel = hotels[i];
                    return Container(
                      margin: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
                      child: Dismissible(
                        direction: DismissDirection.endToStart,
                        background: Container(
                          padding: const EdgeInsets.symmetric(horizontal: 15),
                          alignment: Alignment.centerRight,
                          decoration: BoxDecoration(
                            borderRadius: BorderRadius.circular(10),
                            color: Colors.greenAccent[700],
                          ),
                          child: const Icon(
                            Icons.check,
                            color: Colors.white,
                            size: 30,
                          ),
                        ),
                        key: ValueKey(hotel.id),
                        child: InkWell(
                          onTap: () {
                            Navigator.pushNamed(
                              context,
                              GoogleMapViewHotel.routeName,
                              arguments: {
                                'hotelId': hotel.id,
                                'voyageId': voyage.id
                              },
                            );
                          },
                          child: Card(
                            child: ListTile(
                              title: Text(hotel.name),
                            ),
                          ),
                        ),
                        confirmDismiss: (_) {
                          return Provider.of<VoyageProvider>(context, listen: false)
                              .updateVoyageHotel(voyage, hotel.id)
                              .then((_) => true)
                              .catchError((_) => false);
                        },
                      ),
                    );
                  },
                ),
              ],
            ],
          ),
        );
      },
    );
  }
}