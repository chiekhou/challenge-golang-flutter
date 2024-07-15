import 'package:flutter/material.dart';
import 'package:flutter_app/models/user_model.dart';
import 'package:flutter_app/providers/admin_provider.dart';

class UserManagementScreen extends StatefulWidget {
  static const routeName = '/admin/users';

  const UserManagementScreen({super.key});

  @override
  _UserListWidgetState createState() => _UserListWidgetState();
}

class _UserListWidgetState extends State<UserManagementScreen> {
  List<User> _users = [];
  bool _isLoading = false;

  AdminProvider _apiAdmin = AdminProvider();

  @override
  void initState() {
    super.initState();
    _fetchUsers();
  }

  Future<void> _fetchUsers() async {
    setState(() {
      _isLoading = true;
    });

    try {
      List<User> users = await _apiAdmin.fetchUsers();
      setState(() {
        _users = users;
      });
    } catch (e) {
      print('Error fetching users: $e');
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  Future<void> _deleteUser(int userId) async {
    try {
      bool success = await _apiAdmin.deleteUser(userId);
      if (success) {
        setState(() {
          _users.removeWhere((user) => user.id == userId);
        });
      } else {
        throw Exception('Failed to delete user');
      }
    } catch (e) {
      showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: const Text('Erreur'),
            content: const Text(
                'Échec de la suppression de l\'utilisateur. Veuillez réessayer plus tard.'),
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
      print('Error deleting user: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Liste des Utilisateurs'),
      ),
      body: _isLoading
          ? const Center(child: CircularProgressIndicator())
          : _users.isEmpty
              ? const Center(
                  child: Text('Il n\'y a pas d\'utilisateurs pour l\'instant'),
                )
              : ListView.builder(
                  itemCount: _users.length,
                  itemBuilder: (context, index) {
                    User user = _users[index];
                    return ListTile(
                      title: Text('Votre identifiant: ${user.lastName}'),
                      subtitle: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text('Nom: ${user.lastName}'),
                          Text('Prenom: ${user.firstName}'),
                          Text('Email: ${user.email}'),
                        ],
                      ),
                      trailing: Wrap(
                        spacing: 12,
                        children: <Widget>[
                          IconButton(
                            icon: Icon(Icons.edit),
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
                                        'Êtes-vous sûr de vouloir supprimer cet utilisateur ?'),
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
                                          _deleteUser(user.id);
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
