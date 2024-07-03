import 'package:flutter/material.dart';
import 'package:flutter_app/views/voyage/widgets/voyage_weather.dart.dart';
import 'package:provider/provider.dart';
import '../../providers/destination_provider.dart';
import 'widgets/voyage_activities_hotels.dart';
import 'widgets/voyage_destination_bar.dart';
import '../../models/destination_model.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class VoyageView extends StatelessWidget {
  static const String routeName = '/voyage';

  const VoyageView({super.key});

  @override
  Widget build(BuildContext context) {

    final arguments =
        ModalRoute.of(context)?.settings.arguments as Map<String, dynamic>?;
    if (arguments == null) {

      return Scaffold(
        body: Center(child: Text(AppLocalizations.of(context)!.args_manquants_invalides)),
      );
    }

    if (!arguments.containsKey('destinationName') ||
        arguments['destinationName'] is! String ||
        !arguments.containsKey('voyageId') ||
        arguments['voyageId'] is! int) {

      return Scaffold(
        body: Center(child: Text(AppLocalizations.of(context)!.args_invalides)),
      );
    }

    final String destinationName = arguments['destinationName'] as String;
    final int voyageId = arguments['voyageId'] as int;


    final destinationProvider =
        Provider.of<DestinationProvider>(context, listen: false);
    if (destinationProvider == null) {

      return Scaffold(
        body: Center(child: Text(AppLocalizations.of(context)!.destination_prov_not_found)),
      );
    }

    final Destination destination =
        destinationProvider.getDestinationByName(destinationName);
    if (destination == null) {

      return Scaffold(
        body: Center(child: Text(AppLocalizations.of(context)!.destination_not_found)),
      );
    }

    return Scaffold(
      body: SingleChildScrollView(
        child: Column(
          children: <Widget>[
            VoyageDestinationBar(
              destination: destination,
            ),
            VoyageWeather(destinationName: destinationName),
            VoyageActivitiesHotels(voyageId: voyageId),
          ],
        ),
      ),
    );
  }
}
