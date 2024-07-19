import 'package:flutter_app/models/voyage_model.dart';

import 'member_model.dart';

class Groupe {
  final int id;
  final String nom;
  final double budget;
  final List<dynamic> members;
  final Voyage voyage;
  final int userId;

  Groupe({
    required this.id,
    required this.nom,
    required this.budget,
    required this.members,
    required this.voyage,
    required this.userId,
  });

  factory Groupe.fromJson(Map<String, dynamic> json) {
    return Groupe(
      id: json['id'],
      nom: json['nom'],
      budget: (json['budget'] as num).toDouble(),
      members: (json['members'] as List)
          .map((memberJson) => Member.fromJson(memberJson))
          .toList(),
      voyage: Voyage.fromJson(json['voyage']),
      userId: json['user_id'],
    );
  }
}
