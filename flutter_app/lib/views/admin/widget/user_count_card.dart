import 'package:flutter/material.dart';
import 'package:flutter_app/providers/admin_provider.dart';

class UserCountCard extends StatefulWidget {
  const UserCountCard({super.key});

  @override
  _UserCountCardState createState() => _UserCountCardState();
}

class _UserCountCardState extends State<UserCountCard> {
  late Future<int> _usersSize;
  bool _isLoading = false;

  AdminProvider _apiAdmin = AdminProvider();

  @override
  void initState() {
    super.initState();
    _usersSize = _fetchUsersSize();
  }

  Future<int> _fetchUsersSize() async {
    final users = await _apiAdmin.fetchUsers();
    return users.length;
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<int>(
      future: _fetchUsersSize(),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const Card(
            child: Padding(
              padding: EdgeInsets.all(16.0),
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  Row(
                    children: [
                      Icon(
                        Icons.person,
                        size: 26.0,
                      ),
                      SizedBox(
                        width: 15.0,
                      ),
                      Text(
                        "Utilisateurs",
                        style: TextStyle(
                          fontSize: 26.0,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                    ],
                  ),
                  SizedBox(
                    height: 20.0,
                  ),
                  CircularProgressIndicator(), // Indicateur de chargement
                ],
              ),
            ),
          );
        } else if (snapshot.hasError) {
          return const Card(
            child: Padding(
              padding: EdgeInsets.all(16.0),
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  Row(
                    children: [
                      Icon(
                        Icons.error,
                        size: 26.0,
                        color: Colors.red,
                      ),
                      SizedBox(
                        width: 15.0,
                      ),
                      Text(
                        "Erreur",
                        style: TextStyle(
                          fontSize: 26.0,
                          fontWeight: FontWeight.bold,
                          color: Colors.red,
                        ),
                      ),
                    ],
                  ),
                  SizedBox(
                    height: 20.0,
                  ),
                  Text(
                    'Impossible de charger le nombre d\'utilisateurs',
                    style: TextStyle(
                      fontSize: 18.0,
                      color: Colors.red,
                    ),
                  ),
                ],
              ),
            ),
          );
        } else {
          return Card(
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  const Row(
                    children: [
                      Icon(
                        Icons.person,
                        size: 26.0,
                      ),
                      SizedBox(
                        width: 15.0,
                      ),
                      Text(
                        "Utilisateurs",
                        style: TextStyle(
                          fontSize: 26.0,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                    ],
                  ),
                  const SizedBox(
                    height: 20.0,
                  ),
                  Text(
                    "${snapshot.data} Utilisateurs",
                    style: const TextStyle(
                      fontSize: 36,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ],
              ),
            ),
          );
        }
      },
    );
  }
}
