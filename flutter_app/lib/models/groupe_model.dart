import 'member_model.dart';

class Groupe {
  int id;
  String nom;
  double? budget;
  int userId;
  String? cover;
  List<Member> members;

  Groupe(
      {required this.id,
      required this.nom,
      required this.userId,
      this.cover,
      this.budget,
      required this.members});

  @override
  String toString() {
    return 'Groupe{id: $id, nom: $nom, budget: $budget, userId: $userId, members: $members}';
  }

  factory Groupe.fromJson(Map<String, dynamic> json) {
    var membersFromJson = json['members'] as List;
    List<Member> memberList = membersFromJson
        .map((memberJson) => Member.fromJson(memberJson))
        .toList();
    return Groupe(
      id: json['id'],
      nom: json['nom'] ?? '',
      budget: (json['budget'] as num?)?.toDouble(),
      userId: json['user_id'],
      members: memberList,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'nom': nom,
      'user_id': userId,
      'cover': cover,
      'budget': budget,
      'members': members.map((member) => member.toJson()).toList(),
    };
  }
}
