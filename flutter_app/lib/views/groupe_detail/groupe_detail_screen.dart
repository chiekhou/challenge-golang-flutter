import 'package:flutter/material.dart';
import 'package:flutter_app/views/groupe_detail/widget/group_hotels_activites.dart';
import 'package:provider/provider.dart';
import 'package:web_socket_channel/io.dart';
import '../../models/groupe_model.dart';
import '../../models/member_model.dart';
import '../../providers/auth_provider.dart';
import '../../providers/group_voyage_provider.dart';
import 'widget/groupchat.dart';

class GroupeDetailScreen extends StatelessWidget {
  final int groupeId;
  final TextEditingController emailController = TextEditingController();
  final String host = "10.0.2.2";
  final channel = IOWebSocketChannel.connect('ws://10.0.2.2:8080/ws');

  GroupeDetailScreen({
    super.key,
    required this.groupeId,
  });

  @override
  Widget build(BuildContext context) {
    GroupVoyageProvider groupVoyageProvider = Provider.of<GroupVoyageProvider>(context);
    AuthProvider authProvider = Provider.of<AuthProvider>(context);
    Groupe groupe = groupVoyageProvider.getGroupeById(groupeId);

    return FutureBuilder<Member>(
      future: authProvider.Profile(),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return Center(child: CircularProgressIndicator());
        } else if (snapshot.hasError) {
          return Center(child: Text('Erreur: ${snapshot.error}'));
        } else if (snapshot.hasData) {
          Member user = snapshot.data!;
          return DefaultTabController(
            length: 3,
            child: Scaffold(
              appBar: AppBar(
                title: Text('Détail de groupe'),
                bottom: TabBar(
                  tabs: [
                    Tab(text: groupe.nom),
                    Tab(text: 'Chat'),
                    Tab(text: 'Membres'),
                  ],
                ),
              ),
              body: TabBarView(
                children: [
                  Padding(
                    padding: const EdgeInsets.all(16.0),
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text('Nom: ${groupe.nom}', style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
                        SizedBox(height: 8),
                        Text('Budget: ${groupe.budget?.toString() ?? 'Pas de budget'}'),
                        SizedBox(height: 24),
                        GroupHotelsActivites(
                            groupeId: groupeId,
                            //voyageId: voyageId
                        )
                      ],
                    ),
                  ),
                  GroupChat(
                    groupeId: groupeId,
                    userId: user.id,
                    channel: channel,
                  ),
                  Padding(
                    padding: const EdgeInsets.all(20.0),
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        if(groupe.userId == user.id) ...[
                          SizedBox(height: 32.0),
                          TextFormField(
                            controller: emailController,
                            keyboardType: TextInputType.emailAddress,
                            cursorColor: Colors.blue,
                            textInputAction: TextInputAction.done,
                            decoration: InputDecoration(
                                hintText: 'Invitez un ami en entrant son email'
                            ),
                          ),
                          SizedBox(height: 40.0),
                          ElevatedButton(
                              onPressed: () async {
                                bool success = await groupVoyageProvider.SendInvitation(groupeId, emailController.text);
                                if (success) {
                                  ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('Invitation envoyée!')));
                                } else {
                                  ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text("Échec de l'envoi")));
                                }
                              },
                              child:
                              Text('Envoyer une invitation')
                          ),
                        ],

                        Expanded(
                          child: groupe.members.isEmpty
                              ? Center(
                            child: Text('Aucun membre pour le moment'),
                          )
                              : ListView.builder(
                            itemCount: groupe.members.length,
                            itemBuilder: (context, index) {
                              Member member = groupe.members[index];
                              return Card(
                                child: SizedBox(
                                  height: 80,
                                  child: ListTile(
                                    leading: CircleAvatar(
                                      child: Text(member.username[0].toUpperCase()),
                                    ),
                                    title: Text(member.username),
                                    subtitle: Text(member.email),
                                  ),
                                ),
                              );
                            },
                          ),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            ),
          );
        } else {
          return Center(child: Text('Aucun profil trouvé'));
        }
      },
    );
  }
}
