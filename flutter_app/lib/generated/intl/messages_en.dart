// DO NOT EDIT. This is code generated via package:intl/generate_localized.dart
// This is a library that provides messages for a en locale. All the
// messages from the main program should be duplicated here with the same
// function name.

// Ignore issues from commonly used lints in this file.
// ignore_for_file:unnecessary_brace_in_string_interps, unnecessary_new
// ignore_for_file:prefer_single_quotes,comment_references, directives_ordering
// ignore_for_file:annotate_overrides,prefer_generic_function_type_aliases
// ignore_for_file:unused_import, file_names, avoid_escaping_inner_quotes
// ignore_for_file:unnecessary_string_interpolations, unnecessary_string_escapes

import 'package:intl/intl.dart';
import 'package:intl/message_lookup_by_library.dart';

final messages = new MessageLookup();

typedef String MessageIfAbsent(String messageStr, List<dynamic> args);

class MessageLookup extends MessageLookupByLibrary {
  String get localeName => 'en';

  final messages = _notInlinedMessages(_notInlinedMessages);
  static Map<String, Function> _notInlinedMessages(_) => <String, Function>{
        "accueil": MessageLookupByLibrary.simpleMessage("Home"),
        "alert_date": MessageLookupByLibrary.simpleMessage(
            "Please select a departure date first"),
        "alert_function_flipping": MessageLookupByLibrary.simpleMessage(
            "Trip save feature is disabled"),
        "alert_update_function": MessageLookupByLibrary.simpleMessage(
            "Mise à jour de la fonctionnalité"),
        "annulation_voyage": MessageLookupByLibrary.simpleMessage("Cancem"),
        "attention": MessageLookupByLibrary.simpleMessage("Wait !"),
        "attention_content":
            MessageLookupByLibrary.simpleMessage("You have not entered a date"),
        "deconnexion": MessageLookupByLibrary.simpleMessage("Logout"),
        "echec_update_function":
            MessageLookupByLibrary.simpleMessage("Failed to update feature"),
        "empty_destination":
            MessageLookupByLibrary.simpleMessage("No destination provided."),
        "erreur": MessageLookupByLibrary.simpleMessage("Error"),
        "erreur_sauvegarde":
            MessageLookupByLibrary.simpleMessage("Error saving trip"),
        "label_activites":
            MessageLookupByLibrary.simpleMessage("My activities"),
        "label_decouverte": MessageLookupByLibrary.simpleMessage("Discovery"),
        "label_hotels": MessageLookupByLibrary.simpleMessage("My hôtels"),
        "langage": MessageLookupByLibrary.simpleMessage("Langage"),
        "mes_voyages": MessageLookupByLibrary.simpleMessage("My trips"),
        "message_function_activate":
            MessageLookupByLibrary.simpleMessage("Feature enabled"),
        "message_function_desactivate":
            MessageLookupByLibrary.simpleMessage("Faeture Update"),
        "organisation_titre":
            MessageLookupByLibrary.simpleMessage("Travel organization"),
        "sauvegarde": MessageLookupByLibrary.simpleMessage("Save"),
        "sauvegarde_voyage":
            MessageLookupByLibrary.simpleMessage("Do you want to save?"),
        "title_drawer": MessageLookupByLibrary.simpleMessage("TRIP ESGI"),
        "toogle_drawer":
            MessageLookupByLibrary.simpleMessage("Activate/Disable Trip")
      };
}
