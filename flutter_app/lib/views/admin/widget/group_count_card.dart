import 'package:flutter/material.dart';
import 'package:flutter_app/providers/group_voyage_provider.dart';

class GroupeCountCard extends StatefulWidget {
  const GroupeCountCard({super.key});

  @override
  _GroupeCountCardState createState() => _GroupeCountCardState();
}

class _GroupeCountCardState extends State<GroupeCountCard> {
  late Future<int> _groupesSize;
  bool _isLoading = false;

  GroupVoyageProvider _apiAdminGroup = GroupVoyageProvider();

  @override
  void initState() {
    super.initState();
    _groupesSize = _fetchGroupesSize();
  }

  Future<int> _fetchGroupesSize() async {
    final groupes = await _apiAdminGroup.fetchGroupes();
    return groupes.length;
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<int>(
      future: _fetchGroupesSize(),
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
                        Icons.group,
                        size: 26.0,
                      ),
                      SizedBox(
                        width: 15.0,
                      ),
                      Text(
                        "Groupes Voyage",
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
                  CircularProgressIndicator(),
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
                    'Impossible de charger le nombre de groupes de voyage',
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
                        Icons.group,
                        size: 26.0,
                      ),
                      SizedBox(
                        width: 15.0,
                      ),
                      Text(
                        "Groupes Voyage",
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
                    "${snapshot.data} Groupes ",
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
