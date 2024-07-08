import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;
import 'dart:convert';

class GroupChat extends StatefulWidget {
  final int groupeId;

  GroupChat({required this.groupeId});

  @override
  _GroupChatState createState() => _GroupChatState();
}

class _GroupChatState extends State<GroupChat> {
  late WebSocketChannel channel;
  final TextEditingController _controller = TextEditingController();
  final List<Map<String, dynamic>> _messages = [];
  final String host = "10.0.2.2";


  @override
  void initState() {
    super.initState();
    channel = WebSocketChannel.connect(
      Uri.parse('http://$host:8080/ws'),
    );
    channel.stream.listen((message) {
      setState(() {
        _messages.add(jsonDecode(message));
      });
    });
  }

  void _sendMessage() {
    if (_controller.text.isNotEmpty) {
      final message = jsonEncode({
        'group_id': widget.groupeId,
        'user_id': widget,
        'content': _controller.text,
      });
      channel.sink.add(message);
      _controller.clear();
    }
  }

  @override
  void dispose() {
    channel.sink.close(status.goingAway);
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
