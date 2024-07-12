import 'package:flutter/material.dart';
import 'package:flutter_app/models/hotel_model.dart';
import '../../../models/activity_model.dart';
import 'voyage_activity_hotel_list.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class VoyageActivitiesHotels extends StatelessWidget {
  final int voyageId;

  const VoyageActivitiesHotels({super.key, required this.voyageId});

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 2,
      child: Column(
        children: <Widget>[
          Container(
            color: Theme.of(context).primaryColorDark,
            child: TabBar(
              indicatorColor: Colors.blue[100],
              tabs: <Widget>[
                Tab(
                  text: AppLocalizations.of(context)!.statut_in_progress,
                ),
                Tab(
                  text: AppLocalizations.of(context)!.statut_terminer,
                )
              ],
            ),
          ),
          SizedBox(
            height: 600,
            child: TabBarView(
              physics: const NeverScrollableScrollPhysics(),
              children: <Widget>[
                VoyageActivityHotelList(
                  voyageId: voyageId,
                  filter: ActivityStatus.ongoing,
                  filterh: HotelStatus.ongoing,
                ),
                VoyageActivityHotelList(
                  voyageId: voyageId,
                  filter: ActivityStatus.done,
                  filterh: HotelStatus.done,
                ),

              ],
            ),
          )
        ],
      ),
    );
  }
}
