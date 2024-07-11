import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;
import 'dart:convert';

class GroupChat extends StatefulWidget {
  final int groupeId;
  final WebSocketChannel channel;

  GroupChat({required this.groupeId, required this.channel});

  @override
  _GroupChatState createState() => _GroupChatState();
}

class _GroupChatState extends State<GroupChat> {
  final TextEditingController _controller = TextEditingController();
  final List<Map<String, dynamic>> _messages = [];

  @override
  void initState() {
    super.initState();

    // Écoute des messages WebSocket
    widget.channel.stream.listen((message) {
      setState(() {
        _messages.add(jsonDecode(message));
      });
    });
  }

  void _sendMessage() {
    if (_controller.text.isNotEmpty) {
      final message = jsonEncode({
        'groupe_voyage_id': widget.groupeId,
        'user_id': 1,
        'content': _controller.text,
      });
      widget.channel.sink.add(message);
      _controller.clear();
    }
  }

  @override
  void dispose() {
    widget.channel.sink.close(status.goingAway);
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        Expanded(
          child: ListView.builder(
            itemCount: _messages.length,
            itemBuilder: (context, index) {
              final message = _messages[index];
              return ListTile(
                leading: CircleAvatar(),
                title: Text(message['content']),
                subtitle: Text('User ID: ${message['user_id']}'),
              );
            },
          ),
        ),
        Row(
          children: <Widget>[
            Expanded(
              child: TextField(
                controller: _controller,
                decoration: InputDecoration(labelText: 'Envoyer un message'),
              ),
            ),
            IconButton(
              icon: Icon(Icons.send),
              onPressed: _sendMessage,
            ),
          ],
        ),
      ],
    );
  }
}
