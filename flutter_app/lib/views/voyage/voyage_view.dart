import 'package:flutter/material.dart';
import 'package:flutter_app/views/voyage/widgets/voyage_weather.dart.dart';
import 'package:provider/provider.dart';
import '../../providers/destination_provider.dart';
import 'widgets/voyage_activities_hotels.dart';
import 'widgets/voyage_destination_bar.dart';
import '../../models/destination_model.dart';

class VoyageView extends StatelessWidget {
  static const String routeName = '/voyage';

  const VoyageView({super.key});

  @override
  Widget build(BuildContext context) {
    // Extraire les arguments et ajouter des vérifications
    final arguments =
        ModalRoute.of(context)?.settings.arguments as Map<String, dynamic>?;
    if (arguments == null) {
      print('Arguments manquants');
      return Scaffold(
        body: Center(child: Text('Arguments manquants ou invalides')),
      );
    }

    // Vérification des clés et des types des arguments
    if (!arguments.containsKey('destinationName') ||
        arguments['destinationName'] is! String ||
        !arguments.containsKey('voyageId') ||
        arguments['voyageId'] is! int) {
      print('Arguments invalides: $arguments');
      return Scaffold(
        body: Center(child: Text('Arguments invalides')),
      );
    }

    final String destinationName = arguments['destinationName'] as String;
    final int voyageId = arguments['voyageId'] as int;

    // Vérifiez si le Provider est correctement configuré
    final destinationProvider =
        Provider.of<DestinationProvider>(context, listen: false);
    if (destinationProvider == null) {
      print('DestinationProvider non trouvé');
      return Scaffold(
        body: Center(child: Text('DestinationProvider non trouvé')),
      );
    }

    final Destination destination =
        destinationProvider.getDestinationByName(destinationName);
    if (destination == null) {
      print('Destination non trouvée: $destinationName');
      return Scaffold(
        body: Center(child: Text('Destination non trouvée')),
      );
    }

    print('Destination trouvée: $destinationName, voyageId: $voyageId');

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
