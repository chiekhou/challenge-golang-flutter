import 'package:flutter/material.dart';
import '../../../models/activity_model.dart';

class VoyageActivityList extends StatelessWidget {
  final List<Activity> activities;
  final Function deleteVoyageActivity;

  const VoyageActivityList(
      {super.key,
      required this.activities,
      required this.deleteVoyageActivity});

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemBuilder: (context, index) {
        var activity = activities[index];
        return Card(
          child: ListTile(
            leading: CircleAvatar(
              backgroundImage: NetworkImage(activity.image),
            ),
            title: Text(activity.name),
            trailing: IconButton(
              icon: const Icon(
                Icons.delete,
                color: Colors.red,
              ),
              onPressed: () {
                deleteVoyageActivity(activities[index]);
                ScaffoldMessenger.of(context)
                  ..removeCurrentSnackBar()
                  ..showSnackBar(
                    const SnackBar(
                      content: Text('Activitée supprimée'),
                      backgroundColor: Colors.red,
                      duration: Duration(seconds: 1),
                    ),
                  );
              },
            ),
          ),
        );
      },
      itemCount: activities.length,
    );
  }
}
