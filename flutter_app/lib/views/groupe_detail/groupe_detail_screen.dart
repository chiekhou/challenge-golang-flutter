import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/models/groupe_model.dart';
import 'package:flutter_app/views/groupe_detail/widget/groupchat.dart';
import 'package:provider/provider.dart';
import '../../models/member_model.dart';
import '../../providers/group_voyage_provider.dart';

class GroupeDetailScreen extends StatelessWidget {
  final int groupeId;
  final String? token;
  //final int userId;
  final TextEditingController emailController = TextEditingController();

  GroupeDetailScreen({super.key, required this.groupeId, this.token});

  @override
  Widget build(BuildContext context) {
    GroupVoyageProvider groupVoyageProvider = Provider.of<GroupVoyageProvider>(context);
    Groupe groupe = groupVoyageProvider.getGroupeById(groupeId);

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
                ],
              ),
            ),
            // Onglet Chat
            GroupChat(groupeId: groupeId),
            Padding(
              padding: const EdgeInsets.all(20.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  TextFormField(
                    controller: emailController,
                    keyboardType: TextInputType.emailAddress,
                    cursorColor: Colors.blue,
                    textInputAction: TextInputAction.done,
                    decoration: InputDecoration(
                      hintText: 'Invitez un ami en entrant son email'
                    ),
                  ),
                  ElevatedButton(
                      onPressed: ()async{
                      bool success = await groupVoyageProvider.SendInvitation(this.groupeId, emailController.text);
                      if(success){
                        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('Invitation envoyée!')));
                      }else{
                        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text("Echec de l'envoie")));
                      }
                      },
                      child: Text('Envoyer une invitation')
                  ),
                  Text('Membres:',
                      style: TextStyle(
                          fontSize: 18,
                          fontWeight: FontWeight.bold
                      )
                  ),
                  Expanded(
                    child: ListView.builder(
                      itemCount: groupe.members.length,
                      itemBuilder: (context, index) {
                        Member member = groupe.members[index];
                        return Card(
                          child: SizedBox(
                            height: 80,
                            child: ListTile(
                              leading: CircleAvatar(
                          /*      backgroundImage: member.photo != null && member.photo.isNotEmpty
                                    ? NetworkImage(member.photo)
                                    : AssetImage('../../assets/images/img.png'),*/
                              ),
                              title: Text(member.username),
                              subtitle: Text(member.email),
                            ),
                          ),
                         // title: Text('${member.username}'),
                          //subtitle: Text(member.email),
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
  }
}
