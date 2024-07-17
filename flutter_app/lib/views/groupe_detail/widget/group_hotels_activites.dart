import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/models/activity_model.dart';
import 'package:flutter_app/providers/group_voyage_provider.dart';
import 'package:provider/provider.dart';

import '../../../models/groupe_model.dart';
import '../../../models/hotel_model.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

import '../../google_map/google_map_view.dart';

class GroupHotelsActivites extends StatefulWidget {
  final int groupeId;
  final int? voyageId;

  GroupHotelsActivites({
    required this.groupeId,
    this.voyageId,
  });

  @override
  _GroupHotelsActivitesState createState() => _GroupHotelsActivitesState();
}

class _GroupHotelsActivitesState extends State<GroupHotelsActivites> {
  late int groupeId;
  late ActivityStatus filter = ActivityStatus.ongoing;
  late HotelStatus filterh = HotelStatus.ongoing;

  @override
  void initState() {
    super.initState();
    groupeId = widget.groupeId;
  }

  @override
  Widget build(BuildContext context) {
    GroupVoyageProvider groupVoyageProvider = Provider.of<GroupVoyageProvider>(context);
    final Groupe groupe = groupVoyageProvider.getGroupeById(groupeId);

    final List<Hotel> hotels = groupe.voyage.hotels.where((hotel) => hotel.status == filterh).toList();
    final List<Activity> activities = groupe.voyage.activities.where((activity) => activity.status == filter).toList();

    return SingleChildScrollView(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          if (activities.isEmpty && hotels.isEmpty)
            Padding(
              padding: EdgeInsets.all(32.0),
              child: Text('Aucunes activités ou hôtels enregistrés'),
            ),
          if (activities.isNotEmpty) ...[
            Container(
              margin: const EdgeInsets.symmetric(vertical: 10),
              child: Text(
                AppLocalizations.of(context)!.title_activities,
                style: const TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
              ),
            ),
            ListView.builder(
              shrinkWrap: true,
              physics: const NeverScrollableScrollPhysics(),
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
                          //'voyageId': voyage.id
                        },
                      ),
                      child: Card(
                        child: ListTile(
                          title: Text(activity.name),
                        ),
                      ),
                    ),
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
              margin: const EdgeInsets.symmetric(vertical: 10),
              child: Text(
                AppLocalizations.of(context)!.title_hotels,
                style: const TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
              ),
            ),
            ListView.builder(
              shrinkWrap: true,
              physics: const NeverScrollableScrollPhysics(),
              itemCount: hotels.length,
              itemBuilder: (context, i) {
                Hotel hotel = hotels[i];
                return Container(
                  margin: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
                  child: Card(
                    child: ListTile(
                      title: Text(hotel.name),
                    ),
                  ),
                );
              },
            ),
          ],
        ],
      ),
    );
  }
}
