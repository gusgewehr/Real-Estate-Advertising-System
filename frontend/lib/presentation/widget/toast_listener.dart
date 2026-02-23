import 'package:flutter/material.dart';

class ToastListener extends StatefulWidget {
  final List<ValueNotifier<String>> errorNotifiers;
  final List<ValueNotifier<String>> successNotifiers;
  final Widget child;

  const ToastListener({
    super.key,
    this.errorNotifiers = const [],
    this.successNotifiers = const [],
    required this.child,
  });

  @override
  State<ToastListener> createState() => _ToastListenerState();
}

class _ToastListenerState extends State<ToastListener> {
  final Map<ValueNotifier<String>, VoidCallback> _listeners = {};

  @override
  void initState() {
    super.initState();

    for (final notifier in widget.errorNotifiers) {
      void callback() {
        if (notifier.value.isNotEmpty) {
          WidgetsBinding.instance.addPostFrameCallback((_) {
            if (mounted) _show(notifier.value, Colors.red.shade700, Colors.red.shade100);
          });
        }
      }
      _listeners[notifier] = callback;
      notifier.addListener(callback);
    }

    for (final notifier in widget.successNotifiers) {
      void callback() {
        if (notifier.value.isNotEmpty) {
          final message = notifier.value;
          WidgetsBinding.instance.addPostFrameCallback((_) {
            if (mounted) {
              _show(message, Color(0xFF166534), Color(0xFFBBF7D0));
              notifier.value = "";
            }
          });
        }
      }
      _listeners[notifier] = callback;
      notifier.addListener(callback);
    }
  }

  void _show(String message, Color textColor,Color bgColor) {
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(
        content: Text(
          message,
          style:  TextStyle(color: textColor, fontWeight: FontWeight.bold),
        ),
        backgroundColor: bgColor,
        duration: const Duration(seconds: 3),
        behavior: SnackBarBehavior.floating,
      ),
    );
  }

  @override
  void dispose() {
    for (final entry in _listeners.entries) {
      entry.key.removeListener(entry.value);
    }
    super.dispose();
  }

  @override
  Widget build(BuildContext context) => widget.child;
}