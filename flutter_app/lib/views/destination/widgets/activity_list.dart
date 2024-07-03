import 'package:flutter/material.dart';
import '../../../models/activity_model.dart';
import 'activity_card.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class ActivityList extends StatelessWidget {
  final List<Activity> activities;
  final List<Activity> selectedActivities;
  final Function toggleActivity;

  const ActivityList({
    super.key,
    required this.activities,
    required this.selectedActivities,
    required this.toggleActivity,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      mainAxisSize: MainAxisSize.min,
      children: [
        Container(
          padding: EdgeInsets.all(8.0),
          color: Colors.grey[300],
          child: Text(
            AppLocalizations.of(context)!.title_activities,
            style: TextStyle(
              fontSize: 16,
              fontWeight: FontWeight.bold,
            ),
          ),
        ),

        Flexible(
          fit: FlexFit.loose,
          child: GridView.builder(
            shrinkWrap: true,
            physics: NeverScrollableScrollPhysics(),
            gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: 2,
              crossAxisSpacing: 10,
              mainAxisSpacing: 10,
              childAspectRatio: 3 / 2,
            ),
            itemCount: activities.length,
            itemBuilder: (context, index) {
              final activity = activities[index];
              return ActivityCard(
                activity: activity,
                isSelected: selectedActivities.contains(activity),
                toggleActivity: () => toggleActivity(activity),
              );
            },
          ),
        ),
      ],
    );
  }
}

