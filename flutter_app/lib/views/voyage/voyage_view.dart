import 'package:flutter/material.dart';
import 'package:flutter_app/views/voyage/widgets/voyage_weather.dart.dart';
import 'package:provider/provider.dart';
import '../../providers/destination_provider.dart';
import 'widgets/voyage_activities.dart';
import 'widgets/voyage_destination_bar.dart';
import '../../models/destination_model.dart';

class VoyageView extends StatelessWidget {
  static const String routeName = '/voyage';

  const VoyageView({super.key});

  @override
  Widget build(BuildContext context) {
    final String destinationName = (ModalRoute.of(context)!.settings.arguments
        as Map<String, String>)['destinationName']!;
    final String voyageId = (ModalRoute.of(context)!.settings.arguments
        as Map<String, String>)['voyageId']!;
    final Destination destination =
        Provider.of<DestinationProvider>(context, listen: false)
            .getDestinationByName(destinationName);
    return Scaffold(
      body: SingleChildScrollView(
        child: Column(
          children: <Widget>[
            VoyageDestinationBar(
              destination: destination,
            ),
            VoyageWeather(destinationName: destinationName),
            VoyageActivities(voyageId: voyageId)
          ],
        ),
      ),
    );
  }
}
