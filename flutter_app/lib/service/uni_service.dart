import 'dart:developer';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:uni_links/uni_links.dart';
import 'package:flutter_app/service/context_utility.dart';
import 'package:flutter_app/views/groupe_detail/groupe_detail_screen.dart';

class UniService {
  static String _code = '';
  static String get code => _code;
  static bool get hasCode => _code.isNotEmpty;

  static void reset() => _code = '';

  static init() async {
    try {
      final Uri? uri = await getInitialUri();
      if (uri != null) {
        uniHandler(uri);
      }
    } on PlatformException catch (e) {
      log('Failed to receive the token: $e');
    } on FormatException catch (e) {
      log('Wrong Format received: $e');
    }

    uriLinkStream.listen((Uri? uri) {
      if (uri != null) {
        uniHandler(uri);
      }
    }, onError: (error) {
      log('OnUriLinkError: $error');
    });
  }

  static void uniHandler(Uri uri) {
    if (uri.pathSegments.isEmpty || uri.queryParameters.isEmpty) return;

    if (uri.pathSegments.length >= 3 && uri.pathSegments[0] == 'groupes' && uri.pathSegments[2] == 'join') {
      int groupId = int.parse(uri.pathSegments[1]);
      Map<String, String> param = uri.queryParameters;
      String receivedCode = param['token'] ?? '';

      if (receivedCode.isNotEmpty) {
        Navigator.push(
          ContextUtility.context!,
          MaterialPageRoute(
            builder: (_) => GroupeDetailScreen(groupeId: groupId),
          ),
        );
      }
    }
  }
}
