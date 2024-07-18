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
        "alert_update_function":
            MessageLookupByLibrary.simpleMessage("Update feature"),
        "annulation_voyage": MessageLookupByLibrary.simpleMessage("Cancem"),
        "args_invalides":
            MessageLookupByLibrary.simpleMessage("Invalid arguments"),
        "args_manquants_invalides": MessageLookupByLibrary.simpleMessage(
            "Missing or invalid arguments"),
        "attention": MessageLookupByLibrary.simpleMessage("Wait !"),
        "attention_content":
            MessageLookupByLibrary.simpleMessage("You have not entered a date"),
        "aucun_voyage":
            MessageLookupByLibrary.simpleMessage("No travel at the moment"),
        "date_aller": MessageLookupByLibrary.simpleMessage("Date to Go"),
        "date_retour": MessageLookupByLibrary.simpleMessage("Return Date"),
        "deconnexion": MessageLookupByLibrary.simpleMessage("Logout"),
        "delete_activite":
            MessageLookupByLibrary.simpleMessage("Activity deleted"),
        "delete_hotel": MessageLookupByLibrary.simpleMessage("Hotel deleted"),
        "destination_not_found":
            MessageLookupByLibrary.simpleMessage("Destination not found"),
        "destination_prov_not_found": MessageLookupByLibrary.simpleMessage(
            "DestinationProvider not found"),
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
            MessageLookupByLibrary.simpleMessage("Feature Update"),
        "montant_personne":
            MessageLookupByLibrary.simpleMessage("Amount/person"),
        "organisation_titre":
            MessageLookupByLibrary.simpleMessage("Travel organization"),
        "profile": MessageLookupByLibrary.simpleMessage("Profile"),
        "recherche_ville":
            MessageLookupByLibrary.simpleMessage("Search for a city"),
        "resultat": MessageLookupByLibrary.simpleMessage("No result"),
        "sauvegarde": MessageLookupByLibrary.simpleMessage("Save"),
        "sauvegarde_voyage":
            MessageLookupByLibrary.simpleMessage("Do you want to save?"),
        "selected_date": MessageLookupByLibrary.simpleMessage("Select a date"),
        "statut_in_progress":
            MessageLookupByLibrary.simpleMessage("In progress"),
        "statut_terminer": MessageLookupByLibrary.simpleMessage("Finished"),
        "statut_voyage":
            MessageLookupByLibrary.simpleMessage("No journey is over yet"),
        "title_activities":
            MessageLookupByLibrary.simpleMessage("List of activities"),
        "title_drawer": MessageLookupByLibrary.simpleMessage("TRIP ESGI"),
        "title_home": MessageLookupByLibrary.simpleMessage("Esgi-Trip"),
        "title_hotels": MessageLookupByLibrary.simpleMessage("List of Hotels"),
        "toogle_drawer":
            MessageLookupByLibrary.simpleMessage("Activate/Disable Trip"),
        "voyage_futur": MessageLookupByLibrary.simpleMessage("Future"),
        "voyage_passed": MessageLookupByLibrary.simpleMessage("Past")
      };
}
