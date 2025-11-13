package pid

import (
	"machine/usb/descriptor"
)

var hidDescriptor = descriptor.Append([][]byte{
	descriptor.HIDUsagePageGenericDesktop,
	descriptor.HIDUsageDesktopJoystick,
	descriptor.HIDCollectionApplication,
	descriptor.HIDReportID(1),

	// Buttons (4 bits used, packed into a full byte)
	descriptor.HIDUsagePageButton,
	descriptor.HIDUsageMinimum(1),
	descriptor.HIDUsageMaximum(4),
	descriptor.HIDLogicalMinimum(0),
	descriptor.HIDLogicalMaximum(1),
	descriptor.HIDReportSize(1),
	descriptor.HIDReportCount(4),
	descriptor.HIDInputDataVarAbs,
	// Padding to fill the rest of the byte (4 bits)
	descriptor.HIDReportSize(4),
	descriptor.HIDReportCount(1),
	descriptor.HIDInputConstVarAbs,

	// Hat switch (4 bits), followed by 4 bits padding to make a full byte
	descriptor.HIDUsagePageGenericDesktop,
	descriptor.HIDUsageDesktopHatSwitch,
	descriptor.HIDLogicalMinimum(0),
	descriptor.HIDLogicalMaximum(7),
	descriptor.HIDPhysicalMinimum(0),
	descriptor.HIDPhysicalMaximum(315),
	descriptor.HIDUnit(0x14),
	descriptor.HIDReportCount(1),
	descriptor.HIDReportSize(4),
	descriptor.HIDInputDataVarAbs,
	// Padding nibble after hat to complete the byte
	descriptor.HIDReportSize(4),
	descriptor.HIDReportCount(1),
	descriptor.HIDInputConstVarAbs,

	// Axes: X, Y (each 16 bits, signed)
	descriptor.HIDUsagePageGenericDesktop,
	descriptor.HIDUsageDesktopPointer,
	descriptor.HIDLogicalMinimum(-32767),
	descriptor.HIDLogicalMaximum(32767),
	descriptor.HIDReportCount(2),
	descriptor.HIDReportSize(16),
	descriptor.HIDCollectionPhysical,
	descriptor.HIDUsageDesktopX,
	descriptor.HIDUsageDesktopY,
	descriptor.HIDInputDataVarAbs,
	descriptor.HIDCollectionEnd,
	descriptor.HIDCollectionEnd,
})
