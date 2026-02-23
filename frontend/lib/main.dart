import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:frontend/presentation/page/home.dart';

import 'core/di.dart';
import 'l10n/app_localizations.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();

  const String forcedLocale = String.fromEnvironment('LOCALE', defaultValue: 'pt');

  setupDependencies();

  runApp(MyApp(forcedLocale: forcedLocale));
}

class MyApp extends StatelessWidget {
  final String forcedLocale;
  const MyApp({super.key, required this.forcedLocale});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    final darkBlue = const Color(0xFF1A2B47);

    return MaterialApp(
      locale: Locale(forcedLocale),
      title: 'Flutter Demo',
      localizationsDelegates: [
        AppLocalizations.delegate,
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
      ],
      supportedLocales: [
        Locale('en'),
        Locale('pt'),
      ],
      theme: ThemeData(
        useMaterial3: true,
        fontFamily: 'Roboto',
        colorScheme: ColorScheme.fromSeed(
          seedColor: darkBlue,
        ),
        scaffoldBackgroundColor: Colors.white,
        iconTheme: IconThemeData(
          color: darkBlue, // Cor padrão para todos os ícones
          size: 24.0,              // Tamanho padrão
          opacity: 1.0,            // Opacidade (1.0 é totalmente visível)
        ),
      ),
      home: const Home(),
    );
  }
}

