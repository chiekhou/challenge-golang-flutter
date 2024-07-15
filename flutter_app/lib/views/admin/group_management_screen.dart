import 'package:flutter/material.dart';
import 'package:flutter_app/models/groupe_model.dart';
import 'package:flutter_app/providers/group_voyage_provider.dart';

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
      // Afficher un message d'erreur à l'utilisateur
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur lors de la récupération des groupes')),
      );
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
        title: const Text('Liste des Groupes'),
      ),
      body: _isLoading
          ? const Center(child: CircularProgressIndicator())
          : _groupes.isEmpty
              ? const Center(
                  child: Text('Il n\'y a pas de groupe pour l\'instant'),
                )
              : ListView.builder(
                  itemCount: _groupes.length,
                  itemBuilder: (context, index) {
                    Groupe groupe = _groupes[index];
                    return ListTile(
                      title: Text(groupe.nom),
                      subtitle: Text("${groupe.budget}"),
                      trailing: Wrap(
                        spacing: 12, // Espace entre les boutons
                        children: <Widget>[
                          IconButton(
                            icon: const Icon(Icons.edit),
                            onPressed: () {
                              // Navigation vers l'écran de mise à jour de l'utilisateur
                              // (à implémenter selon tes besoins)
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
                                    title:
                                        const Text('Confirmer la suppression'),
                                    content: const Text(
                                        'Êtes-vous sûr de vouloir supprimer ce groupe ?'),
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
                      ),
                    );
                  },
                ),
    );
  }
}
