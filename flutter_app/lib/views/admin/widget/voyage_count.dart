import 'package:flutter/material.dart';
import 'package:flutter_app/providers/voyage_provider.dart';

class VoyageCountCard extends StatefulWidget {
  const VoyageCountCard({super.key});

  @override
  _VoyageCountCardState createState() => _VoyageCountCardState();
}

class _VoyageCountCardState extends State<VoyageCountCard> {
  late Future<int> _voyagesSize;
  bool _isLoading = false;

  VoyageProvider _apiAdminVoyage = VoyageProvider();

  @override
  void initState() {
    super.initState();
    _voyagesSize = _fetchVoyagesSize();
  }

  Future<int> _fetchVoyagesSize() async {
    final users = await _apiAdminVoyage.fetchData();
    return users.length;
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<int>(
      future: _fetchVoyagesSize(),
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
                        "Voyages Voyage",
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
                    'Impossible de charger le nombre de Voyages de voyage',
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
                        Icons.flight,
                        size: 26.0,
                      ),
                      SizedBox(
                        width: 15.0,
                      ),
                      Text(
                        "Voyages",
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
                    "${snapshot.data} Voyages",
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
