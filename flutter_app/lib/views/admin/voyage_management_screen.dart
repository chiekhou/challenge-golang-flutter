import 'package:flutter/material.dart';
import 'package:flutter_app/providers/destination_provider.dart';
import 'package:flutter_app/providers/voyage_provider.dart';
import 'package:intl/intl.dart';
import 'package:provider/provider.dart';

class VoyageManagementScreen extends StatefulWidget {
  static const routeName = '/admin/voyages';

  const VoyageManagementScreen({super.key});

  @override
  _VoyageManagementScreenState createState() => _VoyageManagementScreenState();
}

class _VoyageManagementScreenState extends State<VoyageManagementScreen> {
  @override
  void initState() {
    super.initState();
    _fetchData();
  }

  Future<void> _fetchData() async {
    await Provider.of<VoyageProvider>(context, listen: false).fetchData();
  }

  Future<void> _deleteVoyage(int voyageId) async {
    try {
      bool success = await Provider.of<VoyageProvider>(context, listen: false)
          .deleteVoyage(voyageId);
      if (success) {
        setState(() {
          Provider.of<VoyageProvider>(context, listen: false)
              .deleteVoyage(voyageId);
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
    final voyageProvider = Provider.of<VoyageProvider>(context);
    final DateFormat formatter = DateFormat('dd/MM/yyyy');

    return Scaffold(
      appBar: AppBar(
        title: const Text('Liste des Voyages'),
      ),
      body: voyageProvider.isLoading
          ? const Center(child: CircularProgressIndicator())
          : voyageProvider.voyages.isEmpty
              ? const Center(
                  child: Text('Il n\'y a pas de voyages pour l\'instant'),
                )
              : Center(
                  child: SingleChildScrollView(
                    scrollDirection: Axis.vertical,
                    child: SizedBox(
                      width: 1000,
                      height: 600,
                      child: DataTable(
                        headingRowColor: MaterialStateProperty.resolveWith(
                            (states) => Colors.grey.shade200),
                        columns: const [
                          DataColumn(label: Text("ID")),
                          DataColumn(label: Text("Destination")),
                          DataColumn(label: Text("Date Aller")),
                          DataColumn(label: Text("Date Retour")),
                          DataColumn(label: Text("Actions")),
                        ],
                        rows: voyageProvider.voyages.map((voyage) {
                          return DataRow(cells: [
                            DataCell(Text(voyage.id.toString())),
                            DataCell(Text(voyage.destination)),
                            DataCell(Text(voyage.dateAller != null
                                ? formatter.format(voyage.dateAller!)
                                : 'N/A')),
                            DataCell(Text(voyage.dateRetour != null
                                ? formatter.format(voyage.dateRetour!)
                                : 'N/A')),
                            DataCell(Row(
                              children: [
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
                                                //   Navigator.of(context).pop();
                                                //   _deleteVoyage(voyage.id);
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
                    ),
                  ),
                ),
    );
  }
}
