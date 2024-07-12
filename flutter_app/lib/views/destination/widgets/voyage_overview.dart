import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'voyage_overview_destination.dart';
import '../../../models/voyage_model.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class VoyageOverview extends StatelessWidget {
  final VoidCallback setDateAller;
  final VoidCallback setDateRetour;
  final Voyage voyage;
  final String destinationName;
  final String destinationImage;
  final double amount;

  const VoyageOverview({
    super.key,
    required this.setDateAller,
    required this.setDateRetour,
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
                    voyage.dateAller != null
                        ? DateFormat("d/M/y").format(voyage.dateAller!)
                        : AppLocalizations.of(context)!.date_aller,
                    style: const TextStyle(fontSize: 15),
                  ),
                ),
                ElevatedButton(
                  onPressed: setDateAller,
                  child: Text(AppLocalizations.of(context)!.selected_date),
                ),

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
                Expanded(
                  child: Text(
                    voyage.dateRetour != null
                        ? DateFormat("d/M/y").format(voyage.dateRetour!)
                        : AppLocalizations.of(context)!.date_retour,
                    style: const TextStyle(fontSize: 15),
                  ),
                ),
                ElevatedButton(
                  onPressed: setDateRetour,
                  child: Text(AppLocalizations.of(context)!.selected_date),
                ),

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
                 Expanded(
                  child: Text(
                    AppLocalizations.of(context)!.montant_personne,
                    style: const TextStyle(fontSize: 15),
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
