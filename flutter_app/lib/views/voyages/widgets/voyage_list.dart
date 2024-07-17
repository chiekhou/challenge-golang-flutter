import 'package:flutter/material.dart';
import 'package:flutter_app/views/voyage/voyage_view.dart';
import 'package:intl/intl.dart';
import '../../../../models/voyage_model.dart';

class VoyageList extends StatelessWidget {
  final List<Voyage> voyages;

  const VoyageList({super.key, required this.voyages});

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemCount: voyages.length,
      itemBuilder: (context, i) {
        var voyage = voyages[i];

        return ListTile(
          title: Text(voyage.destination),
          subtitle: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              voyage.dateAller != null
                  ? Text(
                      "Date Aller : ${DateFormat('d/M/y').format(voyage.dateAller!)}")
                  : const Text('Date Aller : N/A'),
              voyage.dateRetour != null
                  ? Text(
                      "Date Retour : ${DateFormat('d/M/y').format(voyage.dateRetour!)}")
                  : const Text('Date Retour : N/A'),
            ],
          ),
          trailing: const Icon(Icons.info),
          onTap: () => Navigator.pushNamed(
            context,
            VoyageView.routeName,
            arguments: {
              'voyageId': voyage.id,
              'destinationName': voyage.destination,
            },
          ),
        );
      },
    );
  }
}
