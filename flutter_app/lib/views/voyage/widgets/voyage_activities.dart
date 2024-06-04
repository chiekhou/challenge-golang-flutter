import 'package:flutter/material.dart';
import '../../../models/activity_model.dart';
import 'voyage_activity_list.dart';

class VoyageActivities extends StatelessWidget {
  final int voyageId;

  const VoyageActivities({super.key, required this.voyageId});

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 2,
      child: Column(
        children: <Widget>[
          Container(
            color: Theme.of(context).primaryColor,
            child: TabBar(
              indicatorColor: Colors.blue[100],
              tabs: const <Widget>[
                Tab(
                  text: 'En cours',
                ),
                Tab(
                  text: 'Terminées',
                )
              ],
            ),
          ),
          SizedBox(
            height: 600,
            child: TabBarView(
              physics: const NeverScrollableScrollPhysics(),
              children: <Widget>[
                VoyageActivityList(
                  voyageId: voyageId,
                  filter: ActivityStatus.ongoing,
                ),
                VoyageActivityList(
                  voyageId: voyageId,
                  filter: ActivityStatus.done,
                )
              ],
            ),
          )
        ],
      ),
    );
  }
}
