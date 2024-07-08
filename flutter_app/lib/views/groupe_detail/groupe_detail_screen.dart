import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/models/groupe_model.dart';
import 'package:flutter_app/views/groupe_detail/widget/groupchat.dart';
import 'package:provider/provider.dart';
import '../../models/member_model.dart';
import '../../providers/group_voyage_provider.dart';

class GroupeDetailScreen extends StatelessWidget {
  final int groupeId;
  //final int userId;

  const GroupeDetailScreen({super.key, required this.groupeId});

  @override
  Widget build(BuildContext context) {
    GroupVoyageProvider groupVoyageProvider = Provider.of<GroupVoyageProvider>(context);
    Groupe groupe = groupVoyageProvider.getGroupeById(groupeId);

    return DefaultTabController(
      length: 2, // Le nombre d'onglets
      child: Scaffold(
        appBar: AppBar(
          title: Text('Détail de groupe'),
          bottom: TabBar(
            tabs: [
              Tab(text: groupe.nom),
              Tab(text: 'Chat'),
            ],
          ),
        ),
        body: TabBarView(
          children: [
            // Onglet Détails
            Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text('Nom: ${groupe.nom}', style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
                  SizedBox(height: 8),
                  Text('Budget: ${groupe.budget?.toString() ?? 'Pas de budget'}'),
                  SizedBox(height: 16),
                  Text('Membres:', style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
                  Expanded(
                    child: ListView.builder(
                      itemCount: groupe.members.length,
                      itemBuilder: (context, index) {
                        Member member = groupe.members[index];
                        return ListTile(
                          title: Text('${member.first_name} ${member.last_name}'),
                          subtitle: Text(member.email),
                        );
                      },
                    ),
                  ),
                ],
              ),
            ),
            // Onglet Chat
            GroupChat(groupeId: groupeId),
          ],
        ),
      ),
    );
  }
}
