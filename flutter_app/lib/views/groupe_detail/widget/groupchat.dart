import 'dart:async';

import 'package:awesome_notifications/awesome_notifications.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/providers/flipping_provider.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:intl/intl.dart';

class GroupChat extends StatefulWidget {
  final int groupeId;
  final int userId;
  final WebSocketChannel channel;

  GroupChat(
      {required this.groupeId, required this.userId, required this.channel});

  @override
  _GroupChatState createState() => _GroupChatState();
}

class _GroupChatState extends State<GroupChat> {
  final TextEditingController _controller = TextEditingController();
  final List<Map<String, dynamic>> _messages = [];
  final ScrollController _scrollController = ScrollController();
  StreamSubscription? _subscription;
  final StreamController<String> _streamController =
      StreamController<String>.broadcast();

  @override
  void initState() {
    super.initState();
    _fetchPreviousMessages();
    // _setupWebSocketListener();
    _subscribeToStream();
  }

  /* void _setupWebSocketListener() {
    widget.channel.stream.listen((message) {
      _streamController.add(message);
    });
  }*/

  void _subscribeToStream() {
    // Assurez-vous de n'écouter qu'une seule fois
    _subscription = _streamController.stream.listen((message) {
      setState(() {
        _messages.add(jsonDecode(message));
      });
      _scrollToBottom();
    });
  }

  void _sendMessage() {
    if (_controller.text.isNotEmpty) {
      final message = jsonEncode({
        'groupe_voyage_id': widget.groupeId,
        'user_id': widget.userId,
        'content': _controller.text,
      });
      widget.channel.sink.add(message);

      AwesomeNotifications().createNotification(
        content: NotificationContent(
          id: DateTime.now().millisecondsSinceEpoch.remainder(100000),
          channelKey: 'basic_channel',
          title: 'Message du groupe',
          body: _controller.text,
          notificationLayout: NotificationLayout.Default,
        ),
      );

      _controller.clear();
      _scrollToBottom();
    }
  }

  void _fetchPreviousMessages() async {
    final url = isSecure
        ? Uri.https(apiAuthority, '/api/users')
        : Uri.http(apiAuthority, '/api/users');

    final response =
        await http.get(Uri.parse('$url/api/messages/${widget.groupeId}'));
    if (response.statusCode == 200) {
      final List<dynamic> previousMessages = jsonDecode(response.body);
      setState(() {
        _messages.addAll(previousMessages.cast<Map<String, dynamic>>());
      });
      _scrollToBottom();
    }
  }

  void _scrollToBottom() {
    if (_scrollController.hasClients) {
      WidgetsBinding.instance.addPostFrameCallback((_) {
        _scrollController.jumpTo(_scrollController.position.maxScrollExtent);
      });
    }
  }

  @override
  void dispose() {
    _subscription?.cancel();
    _streamController.close();
    widget.channel.sink.close(status.goingAway);
    _scrollController.dispose();
    super.dispose();
  }

  String _formatTimestamp(String timestamp) {
    final dateTime = DateTime.parse(timestamp);
    return DateFormat('HH:mm').format(dateTime);
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        Expanded(
          child: ListView.builder(
            controller: _scrollController,
            itemCount: _messages.length,
            itemBuilder: (context, index) {
              final message = _messages[index];
              final isUserMessage = message['user_id'] == widget.userId;
              return Align(
                alignment: isUserMessage
                    ? Alignment.centerRight
                    : Alignment.centerLeft,
                child: Row(
                  mainAxisSize: MainAxisSize.min,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    if (!isUserMessage) ...[
                      CircleAvatar(
                        backgroundImage: NetworkImage(message['user']['photo']),
                      ),
                      SizedBox(width: 10),
                    ],
                    Flexible(
                      child: Column(
                        crossAxisAlignment: isUserMessage
                            ? CrossAxisAlignment.end
                            : CrossAxisAlignment.start,
                        children: [
                          Container(
                            margin: EdgeInsets.symmetric(
                                vertical: 5, horizontal: 10),
                            padding: EdgeInsets.all(10),
                            decoration: BoxDecoration(
                              color:
                                  isUserMessage ? Colors.purple : Colors.blue,
                              borderRadius: BorderRadius.circular(10),
                            ),
                            child: Text(
                              message['content'],
                              style: TextStyle(color: Colors.white),
                            ),
                          ),
                          SizedBox(height: 5),
                          Row(
                            mainAxisSize: MainAxisSize.min,
                            children: [
                              if (!isUserMessage)
                                Text(
                                  message['user']['username'],
                                  style: TextStyle(
                                      color: Colors.grey[700],
                                      fontWeight: FontWeight.bold),
                                ),
                              SizedBox(width: 5),
                              Text(
                                _formatTimestamp(message['created']),
                                style: TextStyle(
                                    color: Colors.grey[700], fontSize: 10),
                              ),
                            ],
                          ),
                        ],
                      ),
                    ),
                    if (isUserMessage) ...[
                      SizedBox(width: 10),
                      CircleAvatar(
                        backgroundImage: NetworkImage(message['user']['photo']),
                      ),
                    ],
                  ],
                ),
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
