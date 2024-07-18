import 'package:flutter/material.dart';
import 'package:flutter_app/models/destination_model.dart';
import 'package:flutter_app/providers/destination_provider.dart';
import 'package:flutter_app/views/admin/widget/add_destination.dart';
import 'package:flutter_app/views/admin/widget/group_count_card.dart';
import 'package:flutter_app/views/admin/widget/user_count_card.dart';
import 'package:flutter_app/views/admin/widget/voyage_count.dart';
import 'package:provider/provider.dart';

class DashboardHomeScreen extends StatefulWidget {
  static const routeName = '/admin/dashboard/home';
  const DashboardHomeScreen({super.key});

  @override
  State<DashboardHomeScreen> createState() => _DashboardHomeScreenState();
}

class _DashboardHomeScreenState extends State<DashboardHomeScreen> {
  bool isExpanded = false;
  bool _isLoading = false;
  List<Destination> _destinations = [];

  @override
  void initState() {
    super.initState();
    _fetchData();
  }

  Future<void> _fetchData() async {
    setState(() {
      _isLoading = true;
    });

    await Provider.of<DestinationProvider>(context, listen: false).fetchData();
  }

  DestinationProvider _apiDestination = DestinationProvider();

  Future<void> _deleteDestination(int destinationId) async {
    try {
      bool success =
          await Provider.of<DestinationProvider>(context, listen: false)
              .deleteDestination(destinationId);
      if (success) {
        setState(() {
          Provider.of<DestinationProvider>(context, listen: false)
              .deleteDestination(destinationId);
        });
      } else {
        throw Exception('Failed to delete destination');
      }
    } catch (e) {
      showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: const Text('Erreur'),
            content: const Text(
                'Échec de la suppression du voyage. Veuillez réessayer plus tard.'),
            actions: <Widget>[
              TextButton(
                child: const Text('OK'),
                onPressed: () {
                  Navigator.of(context).pop();
                },
              ),
            ],
          );
        },
      );
      print('Error deleting destination: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    final destinationProvider = Provider.of<DestinationProvider>(context);
    return Scaffold(
      body: Row(
        children: [
          Expanded(
            child: Padding(
              padding: const EdgeInsets.all(60.0),
              child: SingleChildScrollView(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.start,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    const SizedBox(
                      height: 20.0,
                    ),
                    const Row(
                      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                      children: [
                        Flexible(
                          child: Padding(
                            padding: EdgeInsets.all(18.0),
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                UserCountCard(),
                              ],
                            ),
                          ),
                        ),
                        SizedBox(width: 16),
                        Flexible(
                          child: Padding(
                            padding: EdgeInsets.all(18.0),
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                GroupeCountCard(),
                              ],
                            ),
                          ),
                        ),
                        SizedBox(width: 16),
                        Flexible(
                          child: Padding(
                            padding: EdgeInsets.all(18.0),
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                VoyageCountCard(),
                              ],
                            ),
                          ),
                        ),
                      ],
                    ),
                    const SizedBox(
                      height: 40.0,
                    ),
                    const Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Column(
                          children: [
                            Text(
                              "Destinations",
                              style: TextStyle(
                                fontWeight: FontWeight.bold,
                                fontSize: 28.0,
                              ),
                            ),
                          ],
                        ),
                      ],
                    ),
                    const SizedBox(
                      height: 40.0,
                    ),
                    Column(
                      crossAxisAlignment: CrossAxisAlignment.stretch,
                      children: [
                        DataTable(
                          headingRowColor: MaterialStateProperty.resolveWith(
                              (states) => Colors.grey.shade200),
                          columns: const [
                            DataColumn(label: Text("ID")),
                            DataColumn(label: Text("Destination")),
                            DataColumn(label: Text("Photo")),
                            DataColumn(label: Text("Actions")),
                          ],
                          rows: destinationProvider.destinations
                              .map((destination) {
                            return DataRow(cells: [
                              DataCell(Text(destination.id.toString())),
                              DataCell(Text(destination.name)),
                              DataCell(Text(destination.image)),
                              DataCell(Row(
                                children: [
                                  IconButton(
                                    icon: const Icon(Icons.add_task),
                                    onPressed: () {
                                      Navigator.pushNamed(context,
                                          AddDestinationForm.routeName);
                                    },
                                  ),
                                  IconButton(
                                    icon: const Icon(Icons.delete),
                                    color: Colors.red,
                                    onPressed: () {
                                      showDialog(
                                        context: context,
                                        builder: (BuildContext context) {
                                          return AlertDialog(
                                            title: const Text(
                                                'Confirmer la suppression'),
                                            content: const Text(
                                                'Êtes-vous sûr de vouloir supprimer ce voyage ?'),
                                            actions: <Widget>[
                                              TextButton(
                                                child: const Text('Annuler'),
                                                onPressed: () {
                                                  Navigator.of(context).pop();
                                                },
                                              ),
                                              TextButton(
                                                child: const Text('Supprimer'),
                                                onPressed: () {
                                                  Navigator.of(context).pop();
                                                  _deleteDestination(
                                                      destination.id);
                                                },
                                              ),
                                            ],
                                          );
                                        },
                                      );
                                    },
                                  ),
                                ],
                              )),
                            ]);
                          }).toList(),
                        ),
                      ],
                    )
                  ],
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
