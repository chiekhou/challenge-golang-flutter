import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/providers/group_voyage_provider.dart';
import 'package:flutter_app/views/add_group/add_group_screen.dart';
import 'package:flutter_app/views/group_voyage/widget/groupeList.dart';
import 'package:flutter_app/widgets/app_drawer.dart';
import 'package:provider/provider.dart';

class GroupevoyageScreen extends StatefulWidget {
  static const String routeName = '/groupes';
  const GroupevoyageScreen({super.key});

  @override
  State<GroupevoyageScreen> createState() => _GroupevoyageState();
}

class _GroupevoyageState extends State<GroupevoyageScreen> {
  late Future<void> _groupsFuture;

  @override
  void initState() {
    super.initState();
    _groupsFuture = Provider.of<GroupVoyageProvider>(context, listen: false).GetGroups();
  }

  @override
  Widget build(BuildContext context) {
    GroupVoyageProvider groupVoyageProvider = Provider.of<GroupVoyageProvider>(context);
    return Scaffold(
      appBar: AppBar(
        title: Text('Mes groupes'),
      ),
      drawer: AppDrawer(),
      floatingActionButton: FloatingActionButton(
        backgroundColor: Colors.purple,
        foregroundColor: Colors.black,
        shape: CircleBorder(),
        onPressed: () {
          Navigator.pushNamed(context, AddGroupScreen.routeName);
        },
        child: Icon(Icons.add),
      ),
      body: FutureBuilder<void>(
        future: _groupsFuture,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(child: Text('Erreur: ${snapshot.error}'));
          } else {
            if (groupVoyageProvider.groupes.isEmpty) {
              return Center(child: Text('Vous n\'avez pas encore de groupe!'));
            } else {
              return Groupelist(
                groupes: groupVoyageProvider.groupes,
              );
            }
          }
        },
      ),
    );
  }
}
