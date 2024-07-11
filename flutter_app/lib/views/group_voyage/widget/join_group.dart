import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:flutter_app/providers/group_voyage_provider.dart';

class JoinGroupScreen extends StatefulWidget {
  final int groupId;
  final String token;

  const JoinGroupScreen({Key? key, required this.groupId, required this.token}) : super(key: key);

  @override
  _JoinGroupScreenState createState() => _JoinGroupScreenState();
}

class _JoinGroupScreenState extends State<JoinGroupScreen> {
  bool _isLoading = false;
  String? _error;

  Future<void> _joinGroup() async {
    setState(() {
      _isLoading = true;
      _error = null;
    });

    try {
      await Provider.of<GroupVoyageProvider>(context, listen: false).JoinGroup(widget.groupId, widget.token);
      Navigator.pushReplacementNamed(context, '/group_detail', arguments: widget.groupId);
    } catch (e) {
      setState(() {
        _error = e.toString();
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  void initState() {
    super.initState();
    _joinGroup();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Rejoindre le Groupe')),
      body: Center(
        child: _isLoading
            ? CircularProgressIndicator()
            : _error != null
            ? Text('Erreur: $_error')
            : Text('Vous avez rejoint le groupe avec succès!'),
      ),
    );
  }
}
