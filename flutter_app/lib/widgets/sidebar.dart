import 'package:flutter/material.dart';

class Sidebar extends StatelessWidget {
  final Function(String) onCategorySelected;

  Sidebar({required this.onCategorySelected});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 150,
      color: Colors.deepPurple.shade400,
      child: ListView(
        children: [
          ListTile(
            iconColor: Colors.white,
            leading: const Icon(Icons.home),
            title: const Text('Home', style: TextStyle(color: Colors.white)),
            onTap: () => onCategorySelected('Home'),
          ),
          ListTile(
            iconColor: Colors.white,
            leading: const Icon(Icons.people),
            title: const Text('Users', style: TextStyle(color: Colors.white)),
            onTap: () => onCategorySelected('Users'),
          ),
          ListTile(
            iconColor: Colors.white,
            leading: const Icon(Icons.groups),
            title: const Text('Groups', style: TextStyle(color: Colors.white)),
            onTap: () => onCategorySelected('Groups'),
          ),
          ListTile(
            iconColor: Colors.white,
            leading: const Icon(Icons.flight),
            title: const Text('Voyages', style: TextStyle(color: Colors.white)),
            onTap: () => onCategorySelected('Trips'),
          ),
        ],
      ),
    );
  }
}
