import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class GroupItems extends StatelessWidget {
  final String nom;
  final String budget;
  final int? nb_members;
  final VoidCallback onTap;
  final Color color;

  const GroupItems({
    required this.nom,
    required this.budget,
    required this.onTap,
    required this.color,
    this.nb_members,
  });

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: onTap,
      child: Container(
        margin: EdgeInsets.symmetric(vertical: 10, horizontal: 15),
        padding: EdgeInsets.all(8),
        decoration: BoxDecoration(
          color: color,
          borderRadius: BorderRadius.circular(15),
          boxShadow: [
            BoxShadow(
              color: Colors.black26,
              blurRadius: 5,
              offset: Offset(0, 2),
            ),
          ],
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: <Widget>[
            Text(
              budget,
              style: TextStyle(color: Colors.white, fontSize: 14),
            ),
            const SizedBox(height: 8),
            Text(
              nom,
              style: TextStyle(color: Colors.white, fontSize: 22, fontWeight: FontWeight.bold),
            ),
            if (nb_members != null) ...[
              const SizedBox(height: 8),
              Text(
                '$nb_members membres',
                style: TextStyle(color: Colors.white70, fontSize: 16),
              ),
            ],
          ],
        ),
      ),
    );
  }
}
