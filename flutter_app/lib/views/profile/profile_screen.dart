import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:flutter_app/providers/auth_provider.dart';

class ProfileScreen extends StatefulWidget {
  static const String routeName = '/profile';
  const ProfileScreen({super.key});

  @override
  _ProfileScreenState createState() => _ProfileScreenState();
}

class _ProfileScreenState extends State<ProfileScreen> {
  late Future<Map<String, dynamic>> _profileFuture;

  @override
  void initState() {
    super.initState();
    // Initialisation de l'appel API pour récupérer le profil
    _profileFuture = Provider.of<AuthProvider>(context, listen: false).Profile();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Profil'),
      ),
      body: FutureBuilder<Map<String, dynamic>>(
        future: _profileFuture,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(child: Text('Erreur: ${snapshot.error}'));
          } else if (snapshot.hasData) {
            var profileData = snapshot.data!['user'];
            return Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text('Nom: ${profileData['last_name']}', style: TextStyle(fontSize: 18)),
                  Text('Prénom: ${profileData['first_name']}', style: TextStyle(fontSize: 18)),
                  Text('Nom: ${profileData['last_name']}', style: TextStyle(fontSize: 18)),
                  Text('Email: ${profileData['email']}', style: TextStyle(fontSize: 18)),
                ],
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
