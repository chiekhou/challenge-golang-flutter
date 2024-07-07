import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/views/groupe_detail/groupe_detail_screen.dart';
import '../../../models/groupe_model.dart';

class Groupelist extends StatelessWidget {
  final List<Groupe> groupes;
  const Groupelist({super.key, required this.groupes});

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemCount: groupes.length,
      itemBuilder: (context, i) {
        var groupe = groupes[i];
        return ListTile(
          title: Text(groupe.nom),
          subtitle: Text(groupe.budget?.toString() ?? 'Pas de budget'),
          onTap: () => Navigator.push(
            context,
            MaterialPageRoute(
                builder: (context)=> GroupeDetailScreen(groupeId: groupe.id!),
            )
          ),
        );
      },
    );
  }
}
