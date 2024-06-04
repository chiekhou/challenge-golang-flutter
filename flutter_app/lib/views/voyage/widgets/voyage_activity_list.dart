import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../../../models/activity_model.dart';
import '../../../models/voyage_model.dart';
import '../../../providers/voyage_provider.dart';
import '../../google_map/google_map_view.dart';

class VoyageActivityList extends StatelessWidget {
  final int voyageId;
  final ActivityStatus filter;

  const VoyageActivityList(
      {super.key, required this.voyageId, required this.filter});

  @override
  Widget build(BuildContext context) {
    return Consumer<VoyageProvider>(
      builder: (context, voyageProvider, child) {
        final Voyage voyage =
            Provider.of<VoyageProvider>(context).getById(voyageId);
        final List<Activity> activities = voyage.activities
            .where((activity) => activity.status == filter)
            .toList();
        return ListView.builder(
          itemCount: activities.length,
          itemBuilder: (context, i) {
            Activity activity = activities[i];
            return Container(
              margin: const EdgeInsets.symmetric(horizontal: 10),
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
                        return Provider.of<VoyageProvider>(context,
                                listen: false)
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
        );
      },
      child: const Text('123'),
    );
  }
}
