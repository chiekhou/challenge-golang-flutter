import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:image_input/image_input.dart';
import 'package:provider/provider.dart';
import 'package:flutter_app/models/member_model.dart'; // Importez le modèle de membre
import 'package:flutter_app/providers/auth_provider.dart';

class ProfileScreen extends StatefulWidget {
  static const String routeName = '/profile';
  const ProfileScreen({super.key});

  @override
  _ProfileScreenState createState() => _ProfileScreenState();
}

class _ProfileScreenState extends State<ProfileScreen> {
  late Future<Member> _profileFuture;

  @override
  void initState() {
    super.initState();
    _profileFuture = Provider.of<AuthProvider>(context, listen: false).Profile();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Profile'),
      ),
      body: FutureBuilder<Member>(
        future: _profileFuture,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(child: Text('Erreur: ${snapshot.error}'));
          } else if (snapshot.hasData) {
            Member profileData = snapshot.data!;
            return Padding(
              padding: const EdgeInsets.all(16.0),
              child: SingleChildScrollView(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Center(
                      child: ProfileAvatar(
                        radius: 50,
                        allowEdit: true,
                        addImageIcon: Container(
                          decoration: BoxDecoration(
                              color: Theme.of(context).colorScheme.primaryContainer,
                              borderRadius: BorderRadius.circular(100)),
                          child: const Padding(
                            padding: EdgeInsets.all(8.0),
                            child: Icon(Icons.add_a_photo),
                          ),
                        ),
                        removeImageIcon: Container(
                          decoration: BoxDecoration(
                              color: Theme.of(context).colorScheme.primaryContainer,
                              borderRadius: BorderRadius.circular(100)),
                          child: const Padding(
                            padding: EdgeInsets.all(8.0),
                            child: Icon(Icons.close),
                          ),
                        ),
                      ),
                    ),
                    SizedBox(height: 20),
                    TextFormField(
                      initialValue: profileData.username,
                      decoration: InputDecoration(
                        labelText: 'Nom d\'utilisateur',
                        border: OutlineInputBorder(),
                      ),
                    ),
                    SizedBox(height: 10),
                    TextFormField(
                      initialValue: profileData.first_name,
                      decoration: InputDecoration(
                        labelText: 'Prénom',
                        border: OutlineInputBorder(),
                      ),
                    ),
                    SizedBox(height: 10),
                    TextFormField(
                      initialValue: profileData.last_name,
                      decoration: InputDecoration(
                        labelText: 'Nom',
                        border: OutlineInputBorder(),
                      ),
                    ),
                    SizedBox(height: 10),
                    TextFormField(
                      initialValue: profileData.email,
                      decoration: InputDecoration(
                        labelText: 'Email',
                        border: OutlineInputBorder(),
                      ),
                    ),
                    SizedBox(height: 20),
                    ElevatedButton(
                      style: ButtonStyle(
                        alignment: Alignment.center,
                      ),
                      onPressed: () {
                        // Logique de mise à jour des coordonnées
                      },
                      child: Text('Mettre à jour mes coordonnées'),
                    ),
                  ],
                ),
              ),
            );
          } else {
            return Center(child: Text('Aucun profil trouvé'));
          }
        },
      ),
    );
  }
}

