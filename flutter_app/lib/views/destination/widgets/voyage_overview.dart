import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'voyage_overview_destination.dart';
import '../../../models/voyage_model.dart';

class VoyageOverview extends StatelessWidget {
  final VoidCallback setDate;
  final Voyage voyage;
  final String destinationName;
  final String destinationImage;
  final double amount;

  const VoyageOverview({
    super.key,
    required this.setDate,
    required this.voyage,
    required this.destinationName,
    required this.amount,
    required this.destinationImage,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          VoyageOverviewDestination(
              destinationName: destinationName,
              destinationImage: destinationImage),
          const SizedBox(
            height: 30,
          ),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 15.0),
            child: Row(
              children: <Widget>[
                Expanded(
                  child: Text(
                    voyage.date != null
                        ? DateFormat("d/M/y").format(voyage.date!)
                        : 'Choisissez une date',
                    style: const TextStyle(fontSize: 15),
                  ),
                ),
                ElevatedButton(
                  onPressed: setDate,
                  child: const Text('Sélectionner une date'),
                )
              ],
            ),
          ),
          const SizedBox(
            height: 30,
          ),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 15.0),
            child: Row(
              children: <Widget>[
                const Expanded(
                  child: Text(
                    'Montant / personne',
                    style: TextStyle(fontSize: 15),
                  ),
                ),
                Text(
                  '$amount \$',
                  style: const TextStyle(
                    fontSize: 15,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ],
            ),
          ),
          const SizedBox(
            height: 30,
          ),
        ],
      ),
    );
  }
}
