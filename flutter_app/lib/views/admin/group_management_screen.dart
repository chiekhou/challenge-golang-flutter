import 'package:flutter/material.dart';
import 'package:flutter_app/models/groupe_model.dart';
import 'package:flutter_app/providers/group_voyage_provider.dart';
import 'package:flutter_app/views/admin/widget/add_group.dart';

class GroupManagementScreen extends StatefulWidget {
  static const routeName = '/admin/groups';

  const GroupManagementScreen({super.key});

  @override
  _GroupListWidgetState createState() => _GroupListWidgetState();
}

class _GroupListWidgetState extends State<GroupManagementScreen> {
  List<Groupe> _groupes = [];
  bool _isLoading = false;

  GroupVoyageProvider _apiAdminGroup = GroupVoyageProvider();

  @override
  void initState() {
    super.initState();
    _fetchGroups();
  }

  Future<void> _fetchGroups() async {
    setState(() {
      _isLoading = true;
    });

    try {
      List<Groupe> groupes = await _apiAdminGroup.fetchGroupes();
      setState(() {
        _groupes = groupes;
      });
    } catch (e) {
      print('Error fetching groups: $e');
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  Future<void> _deleteGroupe(int groupeId) async {
    try {
      bool success = await _apiAdminGroup.deleteGroup(groupeId);
      if (success) {
        setState(() {
          _groupes.removeWhere((groupe) => groupe.id == groupeId);
        });
      } else {
        throw Exception('Failed to delete groupe');
      }
    } catch (e) {
      showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: const Text('Erreur'),
            content: const Text(
                'Échec de la suppression du groupe. Veuillez réessayer plus tard.'),
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
      print('Error deleting groupe: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Liste des Groupes de voyage'),
      ),
      body: _isLoading
          ? const Center(child: CircularProgressIndicator())
          : _groupes.isEmpty
          ? const Center(
        child: Text('Il n\'y a pas de groupe pour l\'instant'),
      )
          : Center(
        child: SingleChildScrollView(
          scrollDirection: Axis.vertical,
          child: SizedBox(
            width: 1000,
            height: 600,
            child: DataTable(
              headingRowColor: WidgetStateProperty.resolveWith(
                      (states) => Colors.grey.shade200),
              columns: const [
                DataColumn(label: Text("ID")),
                DataColumn(label: Text("Nom de groupe")),
                DataColumn(label: Text("Budget Voyage")),
                DataColumn(label: Text("Actions")),
              ],
              rows: _groupes.map((groupe) {
                return DataRow(cells: [
                  DataCell(Text(groupe.id.toString())),
                  DataCell(Text(groupe.nom)),
                  DataCell(Text(groupe.budget.toString())),
                  DataCell(Row(
                    children: [
                      IconButton(
                        icon: Icon(Icons.add_task),
                        onPressed: () {
                          Navigator.pushNamed(
                              context, AddGroupFormAdmin.routeName);
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
                                    'Êtes-vous sûr de vouloir supprimer ce groupe de voyage ?'),
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
                                      _deleteGroupe(groupe.id);
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