import 'voyage_model.dart';

class Groupe {

  final int id;
  final String nom;
  final double budget;
  final int userId;
  final List<dynamic> members;
  final dynamic chats;
  final Voyage voyage;

  Groupe({
    required this.id,
    required this.nom,
    required this.budget,
    required this.userId,
    required this.members,
    this.chats,
    required this.voyage,
  });


  factory Groupe.fromJson(Map<String, dynamic> json) {

    return Groupe(
      id: json['id'],
      nom: json['nom'],
      budget: (json['budget'] as num).toDouble(),
      userId: json['user_id'],
      members: json['members'] ?? [],
      chats: json['chats'],
      voyage: Voyage.fromJson(json['voyage']),
    );
  }
}
