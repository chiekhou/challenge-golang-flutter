import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/views/group_voyage/widget/group_items.dart';
import 'package:flutter_app/views/groupe_detail/groupe_detail_screen.dart';
import 'dart:math';
import '../../../models/groupe_model.dart';

class Groupelist extends StatelessWidget {
  final List<Groupe> groupes;
  const Groupelist({super.key, required this.groupes});

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      scrollDirection: Axis.vertical,
      itemCount: groupes.length,
      itemBuilder: (context, i) {
        var groupe = groupes[i];
        return GroupItems(
          nom: groupe.nom,
          budget: groupe.budget?.toString() ?? 'Pas de budget',
          color: getRandomColor(),
          onTap: () => Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => GroupeDetailScreen(groupeId: groupe.id!),
            ),
          ),
        );
      },
    );
  }

  // Génère une couleur aléatoire
  Color getRandomColor() {
    final Random random = Random();
    return Color.fromARGB(
      255,
      random.nextInt(256),
      random.nextInt(256),
      random.nextInt(256),
    );
  }
}
